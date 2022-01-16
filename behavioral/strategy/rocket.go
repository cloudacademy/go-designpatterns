package main

import "fmt"

type rocket struct {
	mission         string
	maxSpeed        int
	engineStatus    int //0=off, 1=ignited, 2=maxburn, 3=minburn
	landingStrategy iLandingStrategy
	onSurface       bool
}

func buildRocket(mission string, speed int, landing iLandingStrategy) *rocket {
	return &rocket{
		mission:         mission,
		maxSpeed:        speed,
		engineStatus:    0,
		landingStrategy: landing,
		onSurface:       true,
	}
}

func (r *rocket) launch() {
	fmt.Println("launching 3..2..1..")
	r.engineStatus = 1
}

func (r *rocket) fly() {
	fmt.Println("flying...")
	r.engineStatus = 2
	r.onSurface = false
}

func (r *rocket) land() {
	fmt.Println("preparing to land...")
	r.engineStatus = 3
	r.landingStrategy.land(r)
	r.engineStatus = 0
	r.onSurface = true
}
