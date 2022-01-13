package main

import "fmt"

type saturnv struct {
	rocket
}

func (r *saturnv) addFuel(amount int) {
	fmt.Printf("fueling SaturnV rocket with %d litres of fuel...\n", amount)
	r.fuel = amount
}

func (r *saturnv) ignite() bool {
	fmt.Printf("igniting SaturnV rocket...\n")
	return r.fuel > 0
}

func (r *saturnv) throttleUp(speed int) bool {
	fmt.Printf("throttling up SaturnV rocket to %d...\n", speed)
	if speed >= r.minSpeed && speed <= r.maxSpeed {
		return true
	}
	return false
}
