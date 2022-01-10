package main

type saturnvBuilder struct {
	engine     string
	maxspeed   int
	maxpayload int
}

func newSaturnvBuilder() *saturnvBuilder {
	return &saturnvBuilder{}
}

func (b *saturnvBuilder) setEngineType() {
	b.engine = "F1 Rocketdyne"
}

func (b *saturnvBuilder) setMaxSpeed() {
	b.maxspeed = 25000
}

func (b *saturnvBuilder) setMaxPayload() {
	b.maxpayload = 5250000
}

func (b *saturnvBuilder) getRocket() rocket {
	return rocket{
		engine:     b.engine,
		maxspeed:   b.maxspeed,
		maxpayload: b.maxpayload,
	}
}
