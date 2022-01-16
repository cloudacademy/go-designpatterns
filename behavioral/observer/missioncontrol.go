package main

import "fmt"

type missionControl struct {
	name string
}

func (m *missionControl) update(mission string, status string) {
	fmt.Printf("update received, mission control %s - mission %s, status %s\n", m.name, mission, status)
}

func (m *missionControl) getName() string {
	return m.name
}
