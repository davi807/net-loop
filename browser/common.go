package browser

import (
	"log"

	"github.com/zserge/lorca"
)

var ui lorca.UI

// Init initialize browser main ui
func Init(serverRoot string, width int, height int) {
	var err error
	ui, err = lorca.New("http://"+serverRoot, ".cache", width, height)
	if err != nil {
		log.Fatal(err)
	}
}

// Start browser ui and go to serverRoot url
func Start() {
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

// Get returns ui instance for communication
func Get() lorca.UI {
	return ui
}
