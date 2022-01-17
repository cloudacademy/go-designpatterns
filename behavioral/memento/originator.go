package main

type originator struct {
	state string
}

func (org *originator) createMemento() *memento {
	return &memento{state: org.state}
}

func (org *originator) restoreMemento(mem *memento) {
	org.state = mem.getSavedState()
}

func (org *originator) setState(state string) {
	org.state = state
}

func (org *originator) getState() string {
	return org.state
}
