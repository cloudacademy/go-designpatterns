package main

import (
	"fmt"
)

func main() {
	fmt.Print("FLYWEIGHT Example:\n\n")

	rocketSimulation := simulation{
		name: "Launch Test",
	}

	rocketSimulation.load()

	fmt.Printf("Starting %s\n", rocketSimulation.getName())
	rocketSimulation.start()

	fmt.Println()
}
