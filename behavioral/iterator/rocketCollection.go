package main

type rocketCollection struct {
	rockets []*rocket
}

func (coll *rocketCollection) newIterator() iterator {
	return &rocketIterator{
		rockets: coll.rockets,
	}
}
