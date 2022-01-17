package main

import "fmt"

type emptyFuelState struct {
	rocket *rocket
}

func (s *emptyFuelState) fueling(fuel int) error {
	fmt.Println("adding fuel...")
	s.rocket.fuel = fuel
	s.rocket.setState(s.rocket.launchReady)
	return nil
}

func (s *emptyFuelState) ignite() error {
	return fmt.Errorf("fuel tanks empty")
}

func (s *emptyFuelState) increase(speed int) error {
	return fmt.Errorf("engined need to be ignited first")
}
