package main

import "fmt"

func main() {
	fmt.Print("MEDIATOR Example:\n\n")

	missionControl := newMissionControl()

	rocket1 := &falcon9{
		name:     "Mars Lander1",
		mediator: missionControl,
	}

	rocket2 := &falcon9{
		name:     "Mars Lander2",
		mediator: missionControl,
	}

	rocket1.moveToLaunchPad()
	rocket2.moveToLaunchPad()
	rocket1.launch()
	rocket2.launch()

	fmt.Println()
}
