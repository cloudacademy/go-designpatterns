package main

import "fmt"

func main() {
	fmt.Print("TEMPLATEMETHOD Example:\n\n")

	fmt.Println("SaturnV flight plan:")
	saturnv := rocket{
		irocket: &saturnv{minSpeed: 20000, maxSpeed: 25000},
	}
	saturnv.fly(10000, 24500)

	fmt.Println("===================")

	fmt.Println("Falcon9 flight plan:")
	falcon9 := rocket{
		irocket: &falcon9{minSpeed: 20000, maxSpeed: 25000},
	}
	falcon9.fly(5000, 16500)

	fmt.Println()
}
