package main

type istate interface {
	fueling(int) error
	ignite() error
	increase(int) error
}
