package main

type electron struct {
	rocket
}

func newElectron() iRocket {
	stage1 := stage{name: "stage1"}
	stage2 := stage{name: "stage2"}
	stages := []stage{stage1, stage2}

	return &electron{
		rocket: rocket{
			design:   "Electron",
			mission:  "Mars",
			maxspeed: 15000,
			stages:   stages,
		},
	}
}
