package main

import "fmt"

type throttleUpState struct {
	rocket *rocket
}

func (s *throttleUpState) fueling(fuel int) error {
	return fmt.Errorf("rocket already in flight")
}

func (s *throttleUpState) ignite() error {
	return fmt.Errorf("engines already ignited")
}

func (s *throttleUpState) increase(speed int) error {
	if s.rocket.enginesOn {
		if speed >= s.rocket.minSpeed && speed <= s.rocket.maxSpeed {
			s.rocket.currentSpeed = speed
			s.rocket.setState(s.rocket.throttleUp)
			return nil
		}

		return fmt.Errorf("speed would exceed maxspeed")
	}

	return fmt.Errorf("engines need to be ignited first")
}
