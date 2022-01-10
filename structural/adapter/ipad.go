package main

import "fmt"

type ipad struct {
}

func (i *ipad) render() {
	fmt.Println("rendering for **ipad** 2048x2732 display...")
}

func (i *ipad) name() string {
	return "iPad"
}
