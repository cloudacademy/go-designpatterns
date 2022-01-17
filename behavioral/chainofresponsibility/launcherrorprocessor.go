package main

import "fmt"

type launchErrorProcessor struct {
	iprocessor
	next iprocessor
}

func (p *launchErrorProcessor) process(code int, message string) {
	if code == 100 {
		fmt.Printf("system error processor activated (launch): code %d, message [%s]\n", code, message)
	} else if p.next != nil {
		p.next.process(code, message)
	}
}
