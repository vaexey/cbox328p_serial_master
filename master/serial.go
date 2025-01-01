package master

import (
	"bytes"
	"cbox328p/serial_master/panel"
	"errors"
	"fmt"
	"strconv"

	"go.bug.st/serial"
)

func (m *Master) OpenSerial() error {
	if m.port != nil {
		return errors.New("serial port already opened")
	}

	mode := &serial.Mode{
		BaudRate: 38400,
	}

	port, err := serial.Open(m.PortName, mode)

	if err != nil {
		return err
	}

	m.port = port

	return nil
}

func zeroPad(value string, zeros int) string {
	if len(value) >= zeros {
		return value
	}

	var b bytes.Buffer

	for i := len(value); i < zeros; i++ {
		b.WriteByte('0')
	}

	b.WriteString(value)

	return b.String()
}

func (m *Master) SendRaw(command string) error {
	fmt.Printf("TX: %s\n", command)

	if m.port == nil {
		return errors.New("serial port not opened")
	}

	_, err := m.port.Write([]byte(command))

	return err
}

func (m *Master) SendCommand(id string, commandId string, args string) {
	command := fmt.Sprintf("!%s%s%s#", commandId, zeroPad(id, 3), args)

	m.SendRaw(command)
}

func (m *Master) SendDirect(id string, color panel.Color) {
	args := fmt.Sprintf(
		"%s%s%s",
		zeroPad(strconv.Itoa(int(color.R)), 3),
		zeroPad(strconv.Itoa(int(color.G)), 3),
		zeroPad(strconv.Itoa(int(color.B)), 3),
	)

	m.SendCommand(id, "1", args)
}
