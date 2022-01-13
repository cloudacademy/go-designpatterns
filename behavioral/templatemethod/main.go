package main

import "fmt"

func main() {
	fmt.Print("TEMPLATEMETHOD Example:\n\n")

	fmt.Println("SaturnV flight plan:")
	saturnv := rocket{
		rocket: &saturnv{rocket: rocket{minSpeed: 20000, maxSpeed: 25000}},
	}
	saturnv.fly(10000, 24500)

	fmt.Println("===================")

	fmt.Println("Falcon9 flight plan:")
	falcon9 := rocket{
		rocket: &falcon9{rocket: rocket{minSpeed: 15000, maxSpeed: 17000}},
	}
	falcon9.fly(5000, 16500)

	fmt.Println()
}
