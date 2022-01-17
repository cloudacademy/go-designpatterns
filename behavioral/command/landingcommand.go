package main

type launchCommand struct {
	rocket irocket
}

func (c *launchCommand) execute() {
	c.rocket.launch()
}
