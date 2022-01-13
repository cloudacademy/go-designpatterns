package main

import (
	"fmt"
)

type irocket interface {
	addFuel(int)
	ignite() bool
	throttleUp(int) bool
}

type rocket struct {
	rocket   irocket
	fuel     int
	maxSpeed int
	minSpeed int
}

func (r *rocket) fly(fuel int, speed int) error {
	r.rocket.addFuel(fuel)
	if r.rocket.ignite() {
		if r.rocket.throttleUp(speed) {
			return fmt.Errorf("speed wobbles encountered")
		}
	} else {
		return fmt.Errorf("ignition problem encountered")
	}

	return nil
}
