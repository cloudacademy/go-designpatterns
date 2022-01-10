package main

import "fmt"

func main() {
	saturnvBuilder := getBuilder("saturnv")
	electronBuilder := getBuilder("electron")

	director := newDirector(saturnvBuilder)
	saturnv := director.buildRocket()

	fmt.Println("SaturnV:")
	fmt.Printf("Engine Type: %s\n", saturnv.engine)
	fmt.Printf("Max speed: %d\n", saturnv.maxspeed)
	fmt.Printf("Max speed: %d\n", saturnv.maxpayload)

	fmt.Println("================================")

	director.setBuilder(electronBuilder)
	electron := director.buildRocket()

	fmt.Println("Electron:")
	fmt.Printf("Engine Type: %s\n", electron.engine)
	fmt.Printf("Max speed: %d\n", electron.maxspeed)
	fmt.Printf("Max speed: %d\n", electron.maxpayload)
}
