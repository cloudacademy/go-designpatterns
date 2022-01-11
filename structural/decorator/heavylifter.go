package main

type heavylifter struct {
	rocket irocket
}

func (r *heavylifter) getEngines() int {
	engineCount := r.rocket.getEngines()
	return engineCount + 2
}

func (r *heavylifter) getCost() int {
	cost := r.rocket.getCost()
	return cost + 3000000
}
