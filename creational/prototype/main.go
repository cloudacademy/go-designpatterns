package main

import "fmt"

func main() {
	satvStage1 := stage{name: "stage1"}
	satvStage2 := stage{name: "stage2"}
	satvStage3 := stage{name: "stage3"}
	satvStage4 := stage{name: "stage4"}
	satvStages := []stage{satvStage1, satvStage2, satvStage3, satvStage4}
	saturnv_apollo11 := &saturnv{maker: "Nasa", stages: satvStages}
	saturnv_apollo11.setMission("Apollo 11")
	saturnv_apollo12 := saturnv_apollo11.clone()
	saturnv_apollo12.setMission("Apollo 12")

	eStage1 := stage{name: "stage1"}
	eStage2 := stage{name: "stage2"}
	eStages := []stage{eStage1, eStage2}
	electron1 := &electron{maker: "RocketLab", stages: eStages}
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
