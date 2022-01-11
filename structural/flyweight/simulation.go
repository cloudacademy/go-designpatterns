package main

import (
	"math/rand"
)

type simulation struct {
	rockets []*rocket
	name    string
}

func (s *simulation) load() {
	for i := 0; i < 10; i++ {
		if r := rocketFactory("saturnv"); r == nil {
			panic("rocket instance not returned")
		} else {
			s.rockets = append(s.rockets, r)
		}
	}

	for i := 0; i < 10; i++ {
		if r := rocketFactory("falcon9"); r == nil {
			panic("rocket instance not returned")
		} else {
			s.rockets = append(s.rockets, r)
		}
	}
}

func (s *simulation) getName() string {
	return s.name
}

func (s *simulation) start() {
	for _, r := range s.rockets {
		r.launch(rand.Intn(10-3) + 3)
	}
}
