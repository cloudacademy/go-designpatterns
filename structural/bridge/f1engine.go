package main

import "fmt"

type f1engine struct {
}

func (f1engine) ignite() {
	fmt.Println("igniting F1 engine")
}

func (f1engine) throttleUp() {
	fmt.Println("throttling up F1 engine")
}

func (f1engine) throttleDown() {
	fmt.Println("throttling down F1 engine")
}
