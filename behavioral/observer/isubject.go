package main

type isubject interface {
	register(iobserver)
	deregister(iobserver)
	notifyAll(string)
}
