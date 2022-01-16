package main

type rocket interface {
	getType() string
	accept(visitor)
}
