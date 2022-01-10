package main

import "fmt"

func main() {
	fmt.Print("FACTORY Example:\n\n")

	saturnv, _ := getRocket("saturnv")
	electron, _ := getRocket("electron")
	launch(saturnv)
	launch(electron)
}

func launch(rocket iRocket) {
	fmt.Printf("Launching mission: %s\n", rocket.getMission())
	fmt.Printf("Rocket: %s\n", rocket.getDesign())
	fmt.Printf("Maxspeed: %d\n", rocket.getMaxSpeed())
	fmt.Printf("Stages: %d\n", rocket.getStagesCount())
	fmt.Println()
}
