package main

import (
	"log"

	"github.com/zserge/lorca"
)

func startBrowser(serverRoot string) {

	ui, err := lorca.New("http://"+serverRoot, "", 800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
