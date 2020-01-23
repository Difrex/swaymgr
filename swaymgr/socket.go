package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/user"
	"strings"
)

const (
	SOCK_FILE = "/tmp/swaymgr.sock"
)

// ctlCallback controls layout configuration
// available commands is:
// * set: makes layout manageable
//   receives a layout configuation name
//   if name is `manual' this command makes the workspace is
//   non-management
// * get: returns configuration parameters
//   currently, it's only support the `layout' argument
func (m *manager) ctlCallback(callback string) string {
	ctl := strings.Split(strings.Trim(callback, "\n"), " ")
	if len(ctl) < 2 {
		return "error: wrong command length"
	}

	switch ctl[0] {
	case "set":
		if ctl[1] == "manual" {
			if err := m.Unmanage(); err != nil {
				return "error: " + err.Error()
			}
			break
		}
		if _, ok := m.layouts[ctl[1]]; !ok {
			return "warning: unsupported layout " + ctl[1]
		}
		if err := m.layouts[ctl[1]].Manage(); err != nil {
			return fmt.Sprintf("error: %s", err.Error())
		}
	case "get":
		if ctl[1] == "layout" {
			wsConfig := m.getCurrentWorkspaceConfig()
			data, err := json.Marshal(wsConfig)
			if err != nil {
				return "error: " + err.Error()
			}
			return string(data)
		} else {
			return "warning: unsupported command " + ctl[1]
		}
	default:
		return "warning: unsupported command " + ctl[0]
	}

	return "success: ok"
}

// ListenCTL listens a control unix socket
func (m *manager) ListenCTL() {
	ln, err := net.Listen("unix", SOCK_FILE)
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		s, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		callbackResponse := m.ctlCallback(s)

		fmt.Fprintf(conn, "%s\n", callbackResponse)

		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}
}

// SendToCTL sends command from the CLI to the control socket
func SendToCTL(cmd string) {
	conn, err := net.Dial("unix", SOCK_FILE)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, cmd+"\n")
	defer conn.Close()

	s, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print(s)
}

// userID returns user ID string
func userID() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.Uid
}

// cleanUpSocket removes the unix socket file
func cleanUpSocket() error {
	return os.Remove(SOCK_FILE)
}
