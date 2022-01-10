package main

import "fmt"

type saturnv struct {
	maker   string
	mission string
	stages  []stage
}

func (r *saturnv) launch() {
	fmt.Printf("3..2..1 launching %s (%s - %s)\n", "SaturnV", r.maker, r.mission)
}

func (r *saturnv) clone() iRocket {
	//complex cloning logic goes here
	var stages []stage
	stages = append(stages, r.stages...)

	return &saturnv{maker: r.maker, stages: stages}
}

func (r *saturnv) setMission(mission string) {
	r.mission = mission
}
