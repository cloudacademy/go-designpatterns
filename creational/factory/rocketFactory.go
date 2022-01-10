package main

import "fmt"

func getRocket(rocketType string) (iRocket, error) {
	if rocketType == "saturnv" {
		return newSaturnV(), nil
	}
	if rocketType == "electron" {
		return newElectron(), nil
	}
	return nil, fmt.Errorf("wrong rocket type requested")
}
