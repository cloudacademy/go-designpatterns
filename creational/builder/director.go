package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildRocket() rocket {
	d.builder.setEngineType()
	d.builder.setMaxSpeed()
	d.builder.setMaxPayload()
	return d.builder.getRocket()
}
