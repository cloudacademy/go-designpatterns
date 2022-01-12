package main

import "fmt"

func main() {
	fmt.Print("BRIDGE Example:\n\n")

	var f1 engine = &f1engine{}
	var merlin engine = &merlinengine{}

	var saturnv rocket = &saturnv{}

	saturnv.setEngine(f1)
	saturnv.launch()
	fmt.Println()

	saturnv.setEngine(merlin)
	saturnv.launch()
	fmt.Println()

	var falcon9 rocket = &falcon9{}

	falcon9.setEngine(merlin)
	falcon9.launch()
	fmt.Println()

	falcon9.setEngine(f1)
	falcon9.launch()
	fmt.Println()
}
