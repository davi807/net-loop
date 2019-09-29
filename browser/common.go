package browser

import (
	"io/ioutil"
	"log"

	"github.com/zserge/lorca"
)

var ui lorca.UI

// Start initialize browser main ui
func Start(width int, height int) {
	index, err := ioutil.ReadFile("./assets/index.html")
	if err != nil {
		log.Fatal(err)
	}

	ui, err = lorca.New("data:text/html, "+string(index), ".profile", width, height)
	if err != nil {
		log.Fatal(err)
	}

	ui.Bind("load", loader)
}

// Wait for browser ui
func Wait() {
	defer ui.Close()

	// Wait until UI window is closed
	<-ui.Done()
}

// Get returns ui instance for communication
func Get() lorca.UI {
	return ui
}
