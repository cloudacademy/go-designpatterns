package main

import "fmt"

type launchReadyState struct {
	rocket *rocket
}

func (s *launchReadyState) fueling(fuel int) error {
	return fmt.Errorf("rocket already fueled")
}

func (s *launchReadyState) ignite() error {
	if s.rocket.fuel > 0 {
		s.rocket.enginesOn = true
		s.rocket.setState(s.rocket.throttleUp)
	}

	return nil
}

func (s *launchReadyState) increase(speed int) error {
	return fmt.Errorf("engine needs to be ignited first")
}
