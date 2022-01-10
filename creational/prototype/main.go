package main

import "fmt"

func main() {

	saturnv_apollo11 := &saturnv{maker: "Nasa"}
	saturnv_apollo11.setMission("Apollo 11")
	saturnv_apollo12 := saturnv_apollo11.clone()
	saturnv_apollo12.setMission("Apollo 12")

	electron1 := &electron{maker: "RocketLab"}
	electron1.setMission("Moonshot 1")
	electron2 := electron1.clone()
	electron2.setMission("Moonshot 2")

	missions := []iRocket{
		saturnv_apollo11,
		saturnv_apollo12,
		electron1,
		electron2}

	fmt.Println("Launching missions...")
	for _, rocket := range missions {
		rocket.launch()
	}
}
