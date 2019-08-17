package browser

import (
	"log"

	"github.com/zserge/lorca"
)

var ui lorca.UI

// Start browser ui and go to serverRoot url
func Start(serverRoot string, width int, height int) {

	ui, err := lorca.New("http://"+serverRoot, "", width, height)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

// Get returns ui instance for communication
func Get() lorca.UI {
	return ui
}
