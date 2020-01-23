package main

import (
	"flag"
)

var (
	ctlCommand string
)

func init() {
	flag.StringVar(&ctlCommand, "s", "", "Send command to the control socket")
	flag.Parse()
}
