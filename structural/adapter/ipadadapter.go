package main

import "fmt"

type ipadAdapter struct {
	ipadDevice *ipad
}

func (w *ipadAdapter) render() {
	fmt.Println("adapter renders for **ipad** device")
	w.ipadDevice.render()
}

func (w *ipadAdapter) name() string {
	return w.ipadDevice.name()
}
