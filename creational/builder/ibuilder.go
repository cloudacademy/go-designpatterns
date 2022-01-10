package main

type iBuilder interface {
	setEngineType()
	setMaxSpeed()
	setMaxPayload()
	getRocket() rocket
}

func getBuilder(builderType string) iBuilder {
	if builderType == "saturnv" {
		return &saturnvBuilder{}
	}
	if builderType == "electron" {
		return &electronBuilder{}
	}
	return nil
}
