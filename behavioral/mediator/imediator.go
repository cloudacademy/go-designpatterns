package main

type imediator interface {
	canMoveToLaunchPad(irocket) bool
	notifySuccessfulLaunch()
}
