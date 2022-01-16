package main

import (
	"fmt"
)

type totalCost struct {
	cost int
}

func (t *totalCost) visitForSaturnV(r *saturnv) {
	t.cost = 20000000000 + r.fuel*2
	fmt.Printf("Total cost for SaturnV: %d\n", t.cost)
}

func (t *totalCost) visitForFalcon9(r *falcon9) {
	t.cost = r.fuel * 500
	t.cost += r.cargo * 100
	if r.commandcapsule {
		t.cost += 2500000
	}

	fmt.Printf("Total cost for Falcon9: %d\n", t.cost)
}
