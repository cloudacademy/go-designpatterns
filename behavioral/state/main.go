package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("STATE Example:\n\n")

	rocket := newRocket("Mars Lander", 15000, 23500)

	err := rocket.addFuel(1000)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = rocket.igniteEngines()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = rocket.increaseSpeed(15000)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = rocket.increaseSpeed(20000)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
}
