package main

import (
	"fmt"

	"github.com/gballet/hid"
)

func main() {
	fmt.Println(hid.Enumerate(0x2c97, 0x0001))
	fmt.Println(hid.Enumerate(0x534c, 0x0001))
	fmt.Println(hid.Enumerate(0x1209, 0x53c1))
}
