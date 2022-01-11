package main

import (
	"fmt"
)

func main() {
	rocketSimulation := simulation{
		name: "Launch Test",
	}

	rocketSimulation.load()

	fmt.Printf("Starting %s\n", rocketSimulation.getName())
	rocketSimulation.start()
}
