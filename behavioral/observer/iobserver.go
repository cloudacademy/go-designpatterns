package main

type iobserver interface {
	update(string, string)
	getName() string
}
