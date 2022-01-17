package main

import "fmt"

type rocket struct {
	name string

	fuel         int
	minSpeed     int
	maxSpeed     int
	enginesOn    bool
	currentSpeed int

	hasFuel      istate
	launchReady  istate
	throttleUp   istate
	currentState istate
}

func newRocket(name string, minSpeed int, maxSpeed int) *rocket {
	r := &rocket{
		name:      name,
		minSpeed:  minSpeed,
		maxSpeed:  maxSpeed,
		fuel:      0,
		enginesOn: false,
	}

	emptyFuelState := &emptyFuelState{
		rocket: r,
	}
	launchReadyState := &launchReadyState{
		rocket: r,
	}
	throttleUpState := &throttleUpState{
		rocket: r,
	}

	r.setState(emptyFuelState)
	r.hasFuel = emptyFuelState
	r.launchReady = launchReadyState
	r.throttleUp = throttleUpState
	return r
}

func (r *rocket) addFuel(fuel int) error {
	fmt.Println("fueling...")
	return r.currentState.fueling(fuel)
}

func (r *rocket) igniteEngines() error {
	fmt.Println("igniting engines...")
	return r.currentState.ignite()
}

func (r *rocket) increaseSpeed(speed int) error {
	fmt.Println("throttle up...")
	return r.currentState.increase(speed)
}

func (r *rocket) setState(s istate) {
	r.currentState = s
}
