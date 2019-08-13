package main

import "github.com/davi807/net-utils/server"

func main() {

	server.Start()
	startBrowser(server.Root())
}
