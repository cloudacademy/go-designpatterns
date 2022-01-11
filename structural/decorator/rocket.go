package main

type rocket struct {
	irocket
}

func (r *rocket) getEngines() int {
	return 1
}

func (r *rocket) getCost() int {
	return 14000000
}
