package main

import "fmt"

type falcon9 struct {
	enginesOn          bool
	parachutesDeployed bool
}

func (f *falcon9) launch() {
	fmt.Println("blast off")
	f.enginesOn = true
}

func (f *falcon9) land() {
	fmt.Println("engines off")
	f.enginesOn = false
	fmt.Println("parachutes deployed")
	f.parachutesDeployed = true
}
