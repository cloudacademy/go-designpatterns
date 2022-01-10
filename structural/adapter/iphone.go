package main

import "fmt"

type iphone struct {
}

func (i *iphone) render() {
	fmt.Println("rendering for **iphone** 2532x1170 display...")
}

func (i *iphone) name() string {
	return "iPhone"
}
