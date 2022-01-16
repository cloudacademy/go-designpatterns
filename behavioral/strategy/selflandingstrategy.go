package main

import "fmt"

type selflanding struct {
}

func (l *selflanding) land(r *rocket) {
	fmt.Printf("self landing strategy activated (%s)...\n", r.mission)
}
