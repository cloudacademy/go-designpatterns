package main

import "fmt"

type navigationErrorProcessor struct {
	iprocessor
	next iprocessor
}

func (p *navigationErrorProcessor) process(code int, message string) {
	if code == 200 {
		fmt.Printf("system error processor activated (satnav): code %d, message [%s]\n", code, message)
	} else if p.next != nil {
		p.next.process(code, message)
	}
}
