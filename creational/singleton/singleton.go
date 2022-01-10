package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
	//object that should only be created once
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance...")
			singleInstance = &single{}
		} else {
			fmt.Println("check1 - single instance already exists...")
		}
	} else {
		fmt.Println("check2 - single instance already exists...")
	}
	return singleInstance
}
