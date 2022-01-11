package main

import (
	"fmt"
	"sync"
)

var (
	pool map[string]*rocket
	once sync.Once
)

func init() {
	once.Do(func() {
		pool = make(map[string]*rocket)
	})
}

//newRocket - create a new instance
func newRocket(name string) *rocket {
	switch name {
	case "saturnv":
		fmt.Printf("creating saturnv rocket...\n")
		return &rocket{Name: "saturnv", MaxSpeed: 25000}
	case "falcon9":
		fmt.Printf("creating falcon9 rocket...\n")
		return &rocket{Name: "falcon9", MaxSpeed: 20000}
	default:
		fmt.Printf("creating saturnv (default) rocket...\n")
		return &rocket{Name: "saturnv", MaxSpeed: 25000}
	}
}

//rocketFactory - get an instance from pool
func rocketFactory(name string) *rocket {
	if v, ok := pool[name]; ok {
		return v
	}

	rocket := newRocket(name)
	pool[name] = rocket

	return rocket
}
