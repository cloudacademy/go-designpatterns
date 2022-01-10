package main

import "fmt"

type electron struct {
	maker   string
	mission string
}

func (r *electron) launch() {
	fmt.Printf("3..2..1 launching %s (%s - %s)\n", "Electron", r.maker, r.mission)
}

func (r *electron) clone() iRocket {
	//complex cloning logic goes here
	return &electron{maker: r.maker}
}

func (r *electron) setMission(mission string) {
	r.mission = mission
}
