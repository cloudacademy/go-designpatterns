package main

type caretaker struct {
	mementos []*memento
}

func (c *caretaker) addMemento(mem *memento) {
	c.mementos = append(c.mementos, mem)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementos[index]
}
