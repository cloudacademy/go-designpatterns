package main

import "fmt"

type flightErrorProcessor struct {
	iprocessor
	next iprocessor
}

func (p *flightErrorProcessor) process(code int, message string) {
	if code == 300 {
		fmt.Printf("system error processor activated (launch): code %d, message [%s]\n", code, message)
	} else if p.next != nil {
		p.next.process(code, message)
	}
}
