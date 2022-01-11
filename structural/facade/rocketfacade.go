package main

// Facade a high level wrapper of the subsystems
// to provide simple interface to the client
type RocketFacade struct {
	launch *LaunchSystem
	land   *LandingSystem
}

// NewFacade create a new facade
func NewRocketFacade() *RocketFacade {
	return &RocketFacade{
		launch: &LaunchSystem{},
		land:   &LandingSystem{},
	}
}

// Execute a simple interface which provide to client
// to hide the detail usage of the subsystem interfaces
func (f *RocketFacade) Launch() {
	f.launch.Fuel()
	f.launch.Ignite()
	f.launch.ThrottleUp()
}

func (f *RocketFacade) Land() {
	f.land.ThrottleDown()
	f.land.HeatShieldActivated()
	f.land.ParachutesDeployed()
}
