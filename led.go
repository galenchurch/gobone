package gobone

import (
	"fmt"
)

type led struct {
	base, name string
}

var (
	Usr0 = ledBase + "/beaglebone:green:usr0"
	Usr1 = ledBase + "/beaglebone:green:usr1"
	Usr2 = ledBase + "/beaglebone:green:usr2"
	Usr3 = ledBase + "/beaglebone:green:usr3"
)

func LedOff(l string) {
	fmt.Printf("LED %s off\n", l)
	WriteAndClose(l+"/brightness", "0")

}
func LedOn(l string) {
	fmt.Printf("LED %s on\n", l)
	WriteAndClose(l+"/brightness", "1")
}
