package main

type saturnv struct {
	rocket
}

func newSaturnV() iRocket {
	stage1 := stage{name: "stage1"}
	stage2 := stage{name: "stage2"}
	stage3 := stage{name: "stage3"}
	stage4 := stage{name: "stage4"}
	stages := []stage{stage1, stage2, stage3, stage4}

	return &saturnv{
		rocket: rocket{
			design:   "SaturnV",
			mission:  "Apollo11",
			maxspeed: 25000,
			stages:   stages,
		},
	}
}
