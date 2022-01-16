package main

import (
	"fmt"
)

type totalWeight struct {
	weight int
}

func (t *totalWeight) visitForSaturnV(r *saturnv) {
	t.weight = r.fuel + (r.payload + 5000000)
	fmt.Printf("Total weight for SaturnV: %d\n", t.weight)
}

func (t *totalWeight) visitForFalcon9(r *falcon9) {
	t.weight = r.fuel + (r.cargo * 10)
	if r.commandcapsule {
		t.weight += 150000
	}
	fmt.Printf("Total weight for Falcon9: %d\n", t.weight)
}
