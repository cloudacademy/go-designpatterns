package main

type device interface {
	name() string
	render()
}
