package main

import "fmt"

type mission struct {
	isubject
	observers []iobserver
	name      string
}

func newMission(name string) *mission {
	return &mission{
		name: name,
	}
}
func (m *mission) updateStatus(status string) {
	fmt.Printf("updating status: %s...\n", status)
	m.notifyAll(status)
}

func (m *mission) register(o iobserver) {
	m.observers = append(m.observers, o)
}

func (m *mission) deregister(o iobserver) {
	m.observers = removeFromslice(m.observers, o)
}

func (m *mission) notifyAll(status string) {
	for _, observer := range m.observers {
		observer.update(m.name, status)
	}
}

func removeFromslice(observerList []iobserver, observerToRemove iobserver) []iobserver {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getName() == observer.getName() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
