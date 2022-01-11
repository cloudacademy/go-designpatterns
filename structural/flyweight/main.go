package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

//Rocket - struct holds color for reuse
type rocket struct {
	Name     string
	MaxSpeed int
}

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

//rocketFactory - get an instance from the initiaslized pool
func rocketFactory(name string) *rocket {
	if v, ok := pool[name]; ok {
		return v
	}

	rocket := newRocket(name)
	pool[name] = rocket

	return rocket
}

//launch
func (r *rocket) launch(countdown int) {
	fmt.Printf("%s countdown ", r.Name)
	var b strings.Builder
	for i := countdown; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("%d...", i))
	}
	fmt.Println(b.String())
}

type simulation struct {
	rockets []*rocket
	name    string
}

func (s *simulation) load() {
	for i := 0; i < 10; i++ {
		if r := rocketFactory("saturnv"); r == nil {
			panic("rocket instance not returned")
		} else {
			s.rockets = append(s.rockets, r)
		}
	}

	for i := 0; i < 10; i++ {
		if r := rocketFactory("falcon9"); r == nil {
			panic("rocket instance not returned")
		} else {
			s.rockets = append(s.rockets, r)
		}
	}
}

func (s *simulation) getName() string {
	return s.name
}

func (s *simulation) start() {
	for _, r := range s.rockets {
		r.launch(rand.Intn(10-3) + 3)
	}
}

func main() {
	rocketSimulation := simulation{
		name: "Launch Test",
	}

	rocketSimulation.load()

	fmt.Printf("Starting %s\n", rocketSimulation.getName())
	rocketSimulation.start()
}
