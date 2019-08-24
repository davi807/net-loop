package browser

import (
	"io/ioutil"
	"log"

	"github.com/zserge/lorca"
)

var ui lorca.UI

// Init initialize browser main ui
func Init(width int, height int) {
	index, err := ioutil.ReadFile("./assets/index.html")
	if err != nil {
		log.Fatal(err)
	}

	ui, err = lorca.New("data:text/html, "+string(index), ".profile", width, height)
	if err != nil {
		log.Fatal(err)
	}

	bindLoader()
	loadStartContent()
}

// Start browser ui
func Start() {
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

// Get returns ui instance for communication
func Get() lorca.UI {
	return ui
}
