package main

type iRocket interface {
	setEngineType(engine string)
	setMaxSpeed(speed int)
	setMaxPayload(weight int)
	getEngineType() string
	getMaxSpeed() int
	getMaxPayload() int
}

type rocket struct {
	engine     string
	maxspeed   int
	maxpayload int
}

func (r *rocket) setEngineType(engine string) {
	r.engine = engine
}

func (r *rocket) getEngineType() string {
	return r.engine
}

func (r *rocket) setMaxSpeed(speed int) {
	r.maxspeed = speed
}

func (r *rocket) getMaxSpeed() int {
	return r.maxspeed
}

func (r *rocket) setMaxPayload(weight int) {
	r.maxpayload = weight
}

func (r *rocket) getMaxPayload() int {
	return r.maxpayload
}
