package main

type iRocket interface {
	launch()
	setMission(mission string)
	clone() iRocket
}
