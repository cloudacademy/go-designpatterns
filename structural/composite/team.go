package main

import "fmt"

type team struct {
	components []component
	name       string
}

func (t *team) add(c component) {
	t.components = append(t.components, c)
}

func (t *team) getName() string {
	return t.name
}

func (t *team) printDetails() {
	fmt.Printf("Team Name: %s\n", t.getName())
	for _, composite := range t.components {
		composite.printDetails()
	}
}
