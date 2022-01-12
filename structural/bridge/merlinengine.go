package main

import "fmt"

type merlinengine struct {
}

func (merlinengine) ignite() {
	fmt.Println("igniting Merlin engine")
}

func (merlinengine) throttleUp() {
	fmt.Println("throttling up Merlin engine")
}

func (merlinengine) throttleDown() {
	fmt.Println("throttling down Merlin engine")
}
