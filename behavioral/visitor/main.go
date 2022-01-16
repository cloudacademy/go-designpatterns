package main

import "fmt"

func main() {
	fmt.Print("VISITOR Example:\n\n")

	var saturnv rocket = &saturnv{fuel: 1400200, payload: 500000, maxspeed: 25000}
	var falcon9 rocket = &falcon9{fuel: 760030, cargo: 200000, commandcapsule: true}

	totalCost := &totalCost{}

	saturnv.accept(totalCost)
	falcon9.accept(totalCost)

	fmt.Println()

	totalWeight := &totalWeight{}

	saturnv.accept(totalWeight)
	falcon9.accept(totalWeight)

	fmt.Println()
}
