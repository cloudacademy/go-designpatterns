package main

import (
	"fmt"
	"strings"
)

//rocket - struct holds rocket config
type rocket struct {
	Name     string
	MaxSpeed int
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
