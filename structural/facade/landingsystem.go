package main

import "fmt"

type LandingSystem struct{}

func (LandingSystem) ThrottleDown() {
	fmt.Println("landing system - throttling down")
}

func (LandingSystem) HeatShieldActivated() {
	fmt.Println("landing system - heatsheild activated")
}

func (LandingSystem) ParachutesDeployed() {
	fmt.Println("landing system - ignite")
}
