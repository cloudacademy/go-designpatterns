package main

import "fmt"

func main() {
	fmt.Print("OBSERVER Example:\n\n")

	//create new mission
	mission := newMission("Mars Orbiter")

	//create observers
	missionControl1 := &missionControl{name: "Houston"}
	missionControl2 := &missionControl{name: "JPL"}

	//register observers
	mission.register(missionControl1)
	mission.register(missionControl2)

	//mission updates
	mission.updateStatus("launched")
	mission.updateStatus("inflight")
	mission.updateStatus("mars orbit injection")

	fmt.Println()
}
