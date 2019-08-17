package main

import (
	"net-utils/browser"
	"net-utils/server"
)

func main() {
	server.StaticFS("/")
	server.Start()

	browser.Start(server.Root(), 800, 600)
}
