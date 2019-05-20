package main

import (
	"fmt"

	"github.com/karalabe/hid"
)

func main() {
	candidates, err := hid.Enumerate(0x2c97, 0x0001)
	if err == nil {
		for _, dev := range candidates {
			fmt.Println(dev.GetPath())
		}
	}
	candidates, err = hid.Enumerate(0x534c, 0x0001)
	if err == nil {
		for _, dev := range candidates {
			fmt.Println(dev.GetPath())
		}
	}
	candidates, err = hid.Enumerate(0x1209, 0x53c1)
	if err == nil {
		for _, dev := range candidates {
			fmt.Println(dev.GetPath())
		}
	}

}
