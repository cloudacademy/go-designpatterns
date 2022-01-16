package main

type visitor interface {
	visitForSaturnV(*saturnv)
	visitForFalcon9(*falcon9)
}
