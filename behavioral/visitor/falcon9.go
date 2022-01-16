package main

type falcon9 struct {
	fuel           int
	cargo          int
	commandcapsule bool
}

func (f *falcon9) accept(v visitor) {
	v.visitForFalcon9(f)
}

func (f *falcon9) getType() string {
	return "Falcon9"
}
