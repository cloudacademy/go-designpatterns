package main

func main() {
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

	awsdevteam := &team{
		name: "Systems Integration",
	}
	awsdevteam.add(neil)
	awsdevteam.add(devteam)

	awsdevteam.printDetails()
}
