package main

import (
	"fmt"
	"os"

	"github.com/Difrex/gosway/ipc"
)

func main() {
	// CLI implementation.
	// If we receive a client command
	// we need to send it to the control socket
	// and exit from the program.
	if ctlCommand != "" {
		SendToCTL(ctlCommand)
		os.Exit(0)
	}

	// Wait for the syscalls for correct shutdown
	sigWait()

	// Initialize an manager
	manager, err := newManager()
	if err != nil {
		panic(err)
	}
	defer manager.store.dbConn.Close()

	// Listen a control socket
	go manager.ListenCTL()
	defer cleanUpSocket()

	// Subscribe to new Sway events
	o, err := manager.listenerConn.SendCommand(ipc.IPC_SUBSCRIBE, `["window"]`)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(o))

	// Run a events listener
	ch := make(chan *ipc.Event)
	go manager.listenerConn.SubscribeListener(ch)

	// The main loop
	for {
		event := <-ch
		wsConfig, isManaged := manager.isWorkspaceManaged()
		if isManaged {
			switch event.Change {
			case "focus":
				if err := manager.layouts[wsConfig.Layout].OnFocus(event); err != nil {
					fmt.Println("Place error", err)
				}
			case "new":
				if err := manager.layouts[wsConfig.Layout].OnNew(event); err != nil {
					fmt.Println("Place error", err)
				}
			default:
				fmt.Printf("Event received: %s\n", event.Change)
			}
		}
	}
}
