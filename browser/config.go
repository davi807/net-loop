package browser

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const configPath = "./config.json"

type keyValue map[string]string

var config keyValue

func setConfig(name string, value string) {
	config[name] = value
	writeConfig()
}

func writeConfig() {
	data, err := json.Marshal(config)
	if err != nil {
		println(err.Error())
		return
	}

	f, _ := os.Create(configPath)
	f.Write(data)
	f.Close()
}

// RegisterConfigManager js geter/seter for config
func RegisterConfigManager() {
	config = make(keyValue)

	if _, err := os.Stat(configPath); err == nil {
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			println("ERROR: Con't read config file")
			return
		}
		json.Unmarshal(data, &config)
	} else {
		f, _ := os.Create(configPath)
		f.Write([]byte("{}"))
		f.Close()
	}

	ui.Bind("getConfig", func(name string) string { return config[name] })
	ui.Bind("setConfig", setConfig)
	ui.Bind("removeConfig", func(name string) {
		if _, ok := config[name]; !ok {
			return
		}
		delete(config, name)
		writeConfig()
	})
}
