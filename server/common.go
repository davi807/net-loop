package server

import (
	"log"
	"net"
	"net/http"
)

var serverRoot string

func init() {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	serverRoot = l.Addr().String()
	l.Close()
}

// Root returns main server address
func Root() string {
	return serverRoot
}

// Start server
func Start() {
	go func() {
		http.ListenAndServe(serverRoot, nil)
	}()
}
