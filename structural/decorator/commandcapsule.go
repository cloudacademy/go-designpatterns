package main

type commandcapsule struct {
	rocket irocket
}

func (r *commandcapsule) getEngines() int {
	engineCount := r.rocket.getEngines()
	return engineCount + 1
}

func (r *commandcapsule) getCost() int {
	cost := r.rocket.getCost()
	return cost + 8520000
}
