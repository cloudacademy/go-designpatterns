package main

import "fmt"

type falcon9 struct {
	enginesOn          bool
	parachutesDeployed bool
}

func (f *falcon9) launch() {
	fmt.Println("engines on")
	f.enginesOn = true
	fmt.Println("blast off")
}

func (f *falcon9) land() {
	fmt.Println("engines off")
	f.enginesOn = false
	fmt.Println("parachutes deployed")
	f.parachutesDeployed = true
}
