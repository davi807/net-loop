package main

import (
	"net-utils/browser"
)

func main() {
	browser.Init(800, 600)

	browser.RegisterConfigManager()
	browser.SetupTranslation()

	browser.Start()
}
