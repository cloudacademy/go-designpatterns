package main

import "fmt"

type parachutelanding struct {
}

func (l *parachutelanding) land(r *rocket) {
	fmt.Printf("parachute strategy activated (%s)...\n", r.mission)
}
