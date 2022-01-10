package main

import "fmt"

type saturnv struct {
	maker   string
	mission string
}

func (r *saturnv) launch() {
	fmt.Printf("3..2..1 launching %s (%s - %s)\n", "SaturnV", r.maker, r.mission)
}

func (r *saturnv) clone() iRocket {
	//complex cloning logic goes here
	return &saturnv{maker: r.maker}
}

func (r *saturnv) setMission(mission string) {
	r.mission = mission
}
