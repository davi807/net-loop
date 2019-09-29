package workers

import (
	"encoding/json"
	"log"

	"github.com/prologic/bitcask"
)

var devices, statistic *bitcask.Bitcask
var devMeta, statsMeta map[string]string
var metaName = []byte("__meta__")

//OpenDevices opens databeses of devices and statistics
func OpenDevices() {
	var err error
	if devices, err = bitcask.Open("./db/devices"); err != nil {
		log.Fatal(err)
	}
	getMeta(devices, &devMeta)

	if _, ok := devMeta["count"]; !ok {
		devMeta["count"] = "0"
		setMeta(devices, devMeta)
	}
}

// CloseDatabases close all opened databases
func CloseDatabases() {
	if devices != nil {
		devices.Close()
	}

	if statistic != nil {
		statistic.Close()
	}
}

func getMeta(db *bitcask.Bitcask, meta *map[string]string) {
	m := []byte{}

	if *meta == nil {
		*meta = make(map[string]string)
		if db.Has(metaName) {
			m, _ = db.Get(metaName)
		}
	}

	json.Unmarshal(m, meta)
}

func setMeta(db *bitcask.Bitcask, meta map[string]string) {
	metajson, _ := json.Marshal(meta)
	db.Put(metaName, metajson)
}
