package main

import (
	"net-utils/browser"
	"net-utils/workers"
)

func main() {
	go workers.OpenDevices()
	defer workers.CloseDatabases()

	browser.Start(800, 600)

	browser.RegisterDatabaseManager()
	browser.RegisterConfigManager()
	browser.SetupTranslation()

	browser.Wait()
}
