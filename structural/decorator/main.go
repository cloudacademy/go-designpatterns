package main

import "fmt"

func main() {
	fmt.Print("DECORATOR Example:\n\n")

	rocket := &rocket{}

	//add heavy lifting for payload
	fmt.Println("adding heavy lifter feature...")
	heavylifterrocket := &heavylifter{
		rocket: rocket,
	}

	//command capsule
	fmt.Println("adding command capsule feature...")
	commandcapsulerocket := &commandcapsule{
		rocket: heavylifterrocket,
	}

	fmt.Println()

	fmt.Println("Rocket Configuration:")
	fmt.Printf("Engines count: %d\n", commandcapsulerocket.getEngines())
	fmt.Printf("Thrust: %d\n", commandcapsulerocket.getEngines()*15000)
	fmt.Printf("Cost: %d\n", commandcapsulerocket.getCost())

	fmt.Println()
}
