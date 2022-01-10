package main

import "fmt"

func main() {
	fmt.Print("ADAPTER Example:\n\n")

	app := &app{}

	ipadDevice := &ipad{}
	ipadAdapter := &ipadAdapter{
		ipadDevice: ipadDevice,
	}

	//render app for iPad
	app.render(ipadAdapter)

	fmt.Println("=====================")

	iphoneDevice := &iphone{}
	iphoneAdapter := &iphoneAdapter{
		iphoneDevice: iphoneDevice,
	}

	//render app for iPhone
	app.render(iphoneAdapter)
}
