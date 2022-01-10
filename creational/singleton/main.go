package main

import (
	"fmt"
)

func main() {
	fmt.Print("SINGLETON Example:\n\n")

	for i := 0; i < 100; i++ {
		go getInstance()
	}
	fmt.Println()
}
