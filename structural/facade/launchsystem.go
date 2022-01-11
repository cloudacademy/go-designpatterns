package main

import "fmt"

type LaunchSystem struct{}

func (LaunchSystem) Fuel() {
	fmt.Println("launch system - fueling")
}

func (LaunchSystem) Ignite() {
	fmt.Println("launch system - ignite")
}

func (LaunchSystem) ThrottleUp() {
	fmt.Println("launch system - throttling up")
}
