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
	irocket irocket
}

func (r *rocket) fly(fuel int, speed int) error {
	r.irocket.addFuel(fuel)
	if r.irocket.ignite() {
		if r.irocket.throttleUp(speed) {
			return fmt.Errorf("speed wobbles encountered")
		}
	} else {
		return fmt.Errorf("ignition problem encountered")
	}

	return nil
}
