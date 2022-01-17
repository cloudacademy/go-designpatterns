package main

import "fmt"

var (
	orig *originator = &originator{}

	ctaker *caretaker = &caretaker{
		mementos: make([]*memento, 0),
	}

	count = 0
)

func do(direction string) {
	orig.setState(direction)
	ctaker.addMemento(orig.createMemento())
	count = len(ctaker.mementos)
	fmt.Printf("state %d: %s\n", count, orig.getState())
}

func undo() {
	if count > 1 {
		orig.restoreMemento(ctaker.getMemento(count - 2))
		count--
		fmt.Printf("restored to prev state %d: %s\n", count, orig.getState())
	}
}

func main() {
	fmt.Print("MEMENTO Example:\n\n")

	do("up")
	do("up")
	do("right")
	do("left")
	do("up")
	do("right")
	undo()
	undo()
	undo()
	undo()
	undo()

	fmt.Println()
}
