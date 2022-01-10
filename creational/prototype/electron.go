package main

import "fmt"

type electron struct {
	maker   string
	mission string
	stages  []stage
}

func (r *electron) launch() {
	fmt.Printf("3..2..1 launching %s (%s - %s)\n", "Electron", r.maker, r.mission)
}

func (r *electron) clone() iRocket {
	//complex cloning logic goes here
	var stages []stage
	stages = append(stages, r.stages...)
	return &electron{maker: r.maker, stages: stages}
}

func (r *electron) setMission(mission string) {
	r.mission = mission
}
