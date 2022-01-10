package main

import "fmt"

type app struct {
}

func (a *app) render(device device) {
	fmt.Printf("%s: calling render...\n", device.name())
	device.render()
}
