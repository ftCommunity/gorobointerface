package main

import (
	"fmt"
	"time"

	roboif "github.com/ftCommunity/gorobointerface/pkg/robointerface/api"
)

func main() {
	x := roboif.Robointerface{}
	error := x.InitUSB(roboif.GetDefaultUSBConfig())
	if error != nil {
		panic(error)
	}
	for true {
		fmt.Println(x.GetAV())
		time.Sleep(1 * time.Second)
	}
}
