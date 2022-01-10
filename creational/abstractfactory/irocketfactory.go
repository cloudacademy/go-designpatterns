package main

import "fmt"

type iRocketFactory interface {
	makeRocket() iRocket
}

func getRocketFactory(manufacturer string) (iRocketFactory, error) {
	if manufacturer == "nasa" {
		return &nasa{}, nil
	}
	if manufacturer == "spacex" {
		return &spacex{}, nil
	}
	return nil, fmt.Errorf("rocket manufacture not available")
}
