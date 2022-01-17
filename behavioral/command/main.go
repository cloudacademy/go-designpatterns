package main

import "fmt"

func main() {
	fmt.Print("COMMAND Example:\n\n")

	falcon9 := &falcon9{}

	//=======================
	fmt.Println("launch sequence commencing...")
	launchCommand := &launchCommand{
		rocket: falcon9,
	}

	launchToggle := &toggle{
		command: launchCommand,
	}

	launchToggle.move()

	//=======================
	fmt.Println("landing sequence commencing...")
	landCommand := &landCommand{
		rocket: falcon9,
	}

	landingToggle := &toggle{
		command: landCommand,
	}

	landingToggle.move()

	fmt.Println()
}
