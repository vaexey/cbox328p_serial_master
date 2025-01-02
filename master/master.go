package master

import (
	"cbox328p/serial_master/panel"

	"go.bug.st/serial"
)

type Master struct {
	PortName string        `json:"portName"` // Serial port
	Panels   []panel.Panel `json:"panels"`   // Panel array

	panelIdMap      map[string]*panel.Panel   // id -> Panel map
	panelGroupIdMap map[string][]*panel.Panel // id,gid -> []Panels map

	port serial.Port
}
