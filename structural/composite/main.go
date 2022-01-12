package main

import "fmt"

func main() {
	fmt.Print("COMPOSITE Example:\n\n")

	neil := &employee{name: "Neil Armstrong",
		position: "Go Senior Developer",
		manager:  true}

	buzz := &employee{name: "Buzz Aldridge",
		position: "Go Developer",
		manager:  false}

	michael := &employee{name: "Michael Collins",
		position: "Go Developer",
		manager:  false}

	devteam := &team{
		name: "Developers"}

	devteam.add(buzz)
	devteam.add(michael)

	sysintteam := &team{
		name: "Systems Integration",
	}
	sysintteam.add(neil)
	sysintteam.add(devteam)

	sysintteam.printDetails()

	fmt.Println()
}
