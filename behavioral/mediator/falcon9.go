package main

import "fmt"

type falcon9 struct {
	irocket
	name     string
	mediator imediator
}

func (r *falcon9) moveToLaunchPad() {
	fmt.Printf("%s - moving to launch pad...\n", r.name)
	if r.mediator.canMoveToLaunchPad(r) {
		fmt.Printf("%s - at launch pad...\n", r.name)
	} else {
		fmt.Printf("%s - held in queue for launch\n", r.name)
	}
}

func (r *falcon9) launch() {
	fmt.Printf("%s - blast off...\n", r.name)
	r.mediator.notifySuccessfulLaunch()
}
