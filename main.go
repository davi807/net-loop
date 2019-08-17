package main

import (
	"net-utils/browser"
	"net-utils/server"
)

func main() {
	server.StaticFS("/")
	server.Start()

	browser.Init(server.Root(), 800, 600)

	// todo here is going workers registrations

	browser.Start()
}
