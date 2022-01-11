package main

import "fmt"

func main() {
	fmt.Print("FACADE Example:\n\n")

	rocket := &Rocket{}
	rocket.MissionStart()

	fmt.Println()
}
