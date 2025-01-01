package main

import (
	"cbox328p/serial_master/master"
	"cbox328p/serial_master/panel"
	"fmt"
	"time"

	"github.com/crazy3lf/colorconv"
)

func main() {

	fmt.Println("cbox328p serial master app v.1.0.0")

	var m master.Master

	m.Load("master.json")
	m.Save("master.json")

	err := m.OpenSerial()

	fmt.Println(err)

	for {
		for i := 0; i < 360; i++ {
			r, g, b, _ := colorconv.HSLToRGB(float64(i), 1, 0.5)

			m.SendDirect("0", panel.Color{
				R: float32(r), G: float32(g), B: float32(b),
			})

			time.Sleep(10 * time.Millisecond)
		}
	}
}
