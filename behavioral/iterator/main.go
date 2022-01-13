package main

import "fmt"

func main() {

	saturnv := &rocket{
		name:       "SaturnV",
		engine:     "F1 Rocket-Dyne",
		maxspeed:   25000,
		maxpayload: 5000000,
	}
	electron := &rocket{
		name:       "Electron",
		engine:     "Electron",
		maxspeed:   20000,
		maxpayload: 10000,
	}
	falcon9 := &rocket{
		name:       "Falcon9",
		engine:     "Merlin",
		maxspeed:   23000,
		maxpayload: 100000,
	}

	var rockets collection = &rocketCollection{rockets: []*rocket{saturnv, electron, falcon9}}

	itr := rockets.newIterator()

	for itr.hasNext() {
		rocket := itr.getNext()
		fmt.Printf("Rocket %+v\n", rocket)
	}
}
