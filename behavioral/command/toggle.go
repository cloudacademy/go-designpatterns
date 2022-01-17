package main

type toggle struct {
	command command
}

func (t *toggle) move() {
	t.command.execute()
}
