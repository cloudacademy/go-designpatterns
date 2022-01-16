package main

import "fmt"

func main() {
	fmt.Print("STRATEGY Example:\n\n")

	//Mars Lander1 - self landing strategy
	mission1 := "Mars Lander1"
	fmt.Printf("Mission: %s\n", mission1)
	rocket1 := buildRocket(mission1, 1400200, &selflanding{})
	rocket1.launch()
	rocket1.fly()
	rocket1.land()

	fmt.Println()

	//Mars Lander1 - self landing strategy
	mission2 := "Mars Lander2"
	fmt.Printf("Mission: %s\n", mission2)
	rocket2 := buildRocket(mission2, 241000, &parachutelanding{})
	rocket2.launch()
	rocket2.fly()
	rocket2.land()

	fmt.Println()
}
