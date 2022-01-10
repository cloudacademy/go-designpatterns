package main

import "fmt"

func main() {
	nasaFactory, _ := getRocketFactory("nasa")
	spacexFactory, _ := getRocketFactory("spacex")

	saturnv := nasaFactory.makeRocket()
	falcon9 := spacexFactory.makeRocket()

	fmt.Printf("Launching SaturnV:\n")
	fmt.Printf("3..2..1..\n")
	launchRocket(saturnv)

	fmt.Printf("Launching Falcon9:\n")
	fmt.Printf("3..2..1..\n")
	launchRocket(falcon9)
}

func launchRocket(rocket iRocket) {
	fmt.Printf("Engines: %s\n", rocket.getEngineType())
	fmt.Printf("Max Speed: %d\n", rocket.getMaxSpeed())
	fmt.Printf("Max Payload: %d\n", rocket.getMaxPayload())
	fmt.Println()
}
