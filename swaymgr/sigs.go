package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// sigWait waits a unix SYSCALL
func sigWait() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	go watchSignals(sigs)
}

// watchSignals receives SIGTERM, SIGQUIT and SIGINT syscalls from the channel
// if the signal is received, control socket will be deleted and
// a program will be exited with 0 or 2 exit code.
func watchSignals(sigs chan os.Signal) {
	for {
		sig := <-sigs
		fmt.Println(sig)
		if sig == syscall.SIGINT || sig == syscall.SIGQUIT || sig == syscall.SIGTERM {
			err := cleanUpSocket()
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			os.Exit(0)
		}
	}
}
