package workers

import (
	"net"
)

// Device is type of network members
type Device struct {
	ID    int `json:"-"`
	Name  string
	IP    net.IP
	Mac   net.HardwareAddr
	Ports []int
}

func AddDevice()    {}
func UpdateDevice() {}
func RemoveDevice() {}

func GetDevice()   {}
func ListDevices() {}
