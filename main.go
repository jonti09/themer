package main

import (
	"os"

	"github.com/jonti09/themer/server"
)

func main() {
	if os.Args[1] != "--dev" {
		server.Init()
	}
	server.StartAndListen()
}
