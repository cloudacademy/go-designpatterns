package main

type iBuilder interface {
	setEngineType()
	setMaxSpeed()
	setMaxPayload()
	getRocket() rocket
}

func getBuilder(builderType string) iBuilder {
	if builderType == "saturnv" {
		return newSaturnvBuilder()
	}
	if builderType == "electron" {
		return newElectronBuilder()
	}
	return nil
}
