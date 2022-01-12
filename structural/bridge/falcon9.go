package main

import "fmt"

type falcon9 struct {
	engine engine
}

func (r *falcon9) setEngine(e engine) {
	r.engine = e
}

func (r *falcon9) launch() {
	fmt.Println("Falcon9 launch request...")
	r.engine.ignite()
	r.engine.throttleUp()
	r.engine.throttleDown()
}
