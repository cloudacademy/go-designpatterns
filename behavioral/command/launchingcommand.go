package main

type landCommand struct {
	rocket irocket
}

func (c *landCommand) execute() {
	c.rocket.land()
}
