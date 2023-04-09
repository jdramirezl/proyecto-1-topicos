package resolver

import "time"

type Master struct {
	masterIP string
	uptime   int64
}

var (
	master *Master
)

func NewMaster() *Master {
	return &Master{
		masterIP: "",
		uptime:   0,
	}
}

func (m *Master) SetMaster(newIp string) {
	m.masterIP = newIp
	m.uptime = time.Now().UnixNano()
}

func (m *Master) GetMasterIP() string {
	return m.masterIP
}

func GetMaster() *Master {
	return master
}
