package workers

import (
	"log"

	"github.com/prologic/bitcask"
)

var devices, statistic *bitcask.Bitcask

//OpenDevices opens databeses of devices and statistics
func OpenDevices() {
	var err error
	devices, err = bitcask.Open("./db/devices")
	if err != nil {
		log.Fatal(err)
	}
	statistic, _ = bitcask.Open("./db/statistic")
}

// CloseDatabases close all opened databases
func CloseDatabases() {
	devices.Close()
	statistic.Close()
}
