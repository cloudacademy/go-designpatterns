package main

import "fmt"

type falcon9 struct {
	fuel     int
	maxSpeed int
	minSpeed int
}

func (r *falcon9) addFuel(amount int) {
	fmt.Printf("fueling Falcon9 rocket with %d litres of fuel...\n", amount)
	r.fuel = amount
}

func (r *falcon9) ignite() bool {
	fmt.Printf("igniting Falcon9 rocket...\n")
	return r.fuel > 0
}

func (r *falcon9) throttleUp(speed int) bool {
	fmt.Printf("throttling up Falcon9 rocket to %d...\n", speed)
	if speed >= r.minSpeed && speed <= r.maxSpeed {
		return true
	}
	return false
}
