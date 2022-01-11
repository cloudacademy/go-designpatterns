package main

// Client a client which used to call the interface methods
// of different subsystems, from now on, it just need to
// reference to the facade, and use its simple interface
type Rocket struct{}

// Run call facade's simple interface
func (Rocket) MissionStart() {
	f := NewRocketFacade()
	f.Launch()
	f.Land()
}
