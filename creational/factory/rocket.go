package main

type rocket struct {
	design   string
	mission  string
	maxspeed int
	stages   []stage
}

func (r *rocket) getDesign() string {
	return r.design
}

func (r *rocket) getMission() string {
	return r.mission
}

func (r *rocket) getMaxSpeed() int {
	return r.maxspeed
}

func (r *rocket) getStagesCount() int {
	return len(r.stages)
}
