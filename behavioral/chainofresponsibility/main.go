package main

import "fmt"

func buildChain() iprocessor {
	chain := launchErrorProcessor{next: &navigationErrorProcessor{next: &flightErrorProcessor{}}}
	return &chain
}

func main() {
	fmt.Print("CHAIN OF RESPONSIBILITY Example:\n\n")

	chain := buildChain()
	chain.process(100, "problem with iginition detected")
	chain.process(200, "problem with satnav detected")
	chain.process(300, "problem with flight control detected")

	fmt.Println()
}
