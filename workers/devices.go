package workers

import (
	"encoding/json"
	"net"
	"strconv"
)

// Device is type of network members
type Device struct {
	ID    int `json:"-"`
	Name  string
	IP    net.IP
	Mac   net.HardwareAddr
	Ports []int
}

// GetDevice return device by id
func GetDevice(id string) Device {
	dev, err := devices.Get([]byte(id))

	if err != nil {
		println(err.Error())
		return Device{ID: 0}
	}

	var device Device
	json.Unmarshal(dev, &device)

	device.ID, _ = strconv.Atoi(id)
	return device
}

// ListDevices return all devices
func ListDevices() []Device {
	list := []Device{}

	devices.Fold(func(id []byte) error {
		dev := GetDevice(string(id))
		list = append(list, dev)
		return nil
	})

	return list
}

/* this metods can be directly bind to ui */
/*******************************************/

// RemoveDevice remove from database
func RemoveDevice(id string) bool {
	err := removeFild(devices, id)

	if err != nil {
		println(err.Error())
		return false
	}

	setMeta(devices, devMeta)
	return true
}

// AddDevice put device string to database
func AddDevice(device string) bool {
	addCount(devMeta, 1)

	if err := devices.Put([]byte(devMeta["count"]), []byte(device)); err != nil {
		println(err.Error())
		addCount(devMeta, -1)
		return false
	}

	setMeta(devices, devMeta)
	return true
}

// UpdateDevice remove id from json string add to db
func UpdateDevice(id string, device string) bool {
	place := make(map[string]interface{})

	if err := json.Unmarshal([]byte(device), &place); err != nil {
		println(err.Error())
		return false
	}

	if _, ok := place["ID"]; ok {
		delete(place, "ID")
	}

	data, err := json.Marshal(place)

	if err != nil {
		println(err.Error())
		return false
	}

	devices.Put([]byte(id), data)

	return true
}
