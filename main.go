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
	m.GenerateIndex()

	err := m.OpenSerial()

	fmt.Println(err)

	go updateLoop(&m)

	for {
		for i := 0; i < 360; i++ {
			r, g, b, _ := colorconv.HSLToRGB(float64(i), 1, 0.5)

			m.FindPanel("1").Set(panel.Color{
				R: float32(r) / 255, G: float32(g) / 255, B: float32(b) / 255,
			})

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func updateLoop(m *master.Master) {
	for {
		m.SendDirtyPanels()

		time.Sleep(10 * time.Millisecond)
	}
}
