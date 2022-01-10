package main

type stage struct {
	name string
}

type iRocket interface {
	getDesign() string
	getMission() string
	getMaxSpeed() int
	getStagesCount() int
}
