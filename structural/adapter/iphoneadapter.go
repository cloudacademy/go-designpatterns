package main

import "fmt"

type iphoneAdapter struct {
	iphoneDevice *iphone
}

func (w *iphoneAdapter) render() {
	fmt.Println("adapter renders for **iphone** device")
	w.iphoneDevice.render()
}

func (w *iphoneAdapter) name() string {
	return w.iphoneDevice.name()
}
