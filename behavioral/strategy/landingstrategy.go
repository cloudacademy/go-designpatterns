package main

type iLandingStrategy interface {
	land(r *rocket)
}
