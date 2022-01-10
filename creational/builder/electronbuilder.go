package main

type electronBuilder struct {
	engine     string
	maxspeed   int
	maxpayload int
}

func newElectronBuilder() *electronBuilder {
	return &electronBuilder{}
}

func (b *electronBuilder) setEngineType() {
	b.engine = "Rutherford"
}

func (b *electronBuilder) setMaxSpeed() {
	b.maxspeed = 20000
}

func (b *electronBuilder) setMaxPayload() {
	b.maxpayload = 225
}

func (b *electronBuilder) getRocket() rocket {
	return rocket{
		engine:     b.engine,
		maxspeed:   b.maxspeed,
		maxpayload: b.maxpayload,
	}
}
