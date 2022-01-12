package main

import "fmt"

type saturnv struct {
	engine engine
}

func (r *saturnv) setEngine(e engine) {
	r.engine = e
}

func (r *saturnv) launch() {
	fmt.Println("SaturnV launch request...")
	r.engine.ignite()
	r.engine.throttleUp()
	r.engine.throttleDown()
}
