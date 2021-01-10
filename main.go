package main

import (
	"os"

	"github.com/jonti09/themer/server"
)

func hasDevFlag() bool {
	return len(os.Args) == 2 && os.Args[1] == "--dev"
}

func main() {
	if !hasDevFlag() {
		server.Init()
	}
	server.StartAndListen()
}
