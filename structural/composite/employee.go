package main

import "fmt"

type employee struct {
	name     string
	position string
	manager  bool
}

func (e *employee) getName() string {
	return e.name
}

func (e *employee) getPosition() string {
	return e.position
}

func (e *employee) isManager() bool {
	return e.manager
}

func (e *employee) printDetails() {
	fmt.Printf("Name: %s, Position: %s, Manager: %t\n", e.getName(), e.getPosition(), e.isManager())
}
