package gobone

import (
	"fmt"
)

type LED struct {
	Sys   sysfs_led
	Value string //"0" or "1"
}

//LED Globals
var (
	Usr0 = LED{Sys: sysfs_led{SysfsLOC: "/beaglebone:green:usr0", Name: "usr0"}}
	Usr1 = LED{Sys: sysfs_led{SysfsLOC: "/beaglebone:green:usr1", Name: "usr1"}}
	Usr2 = LED{Sys: sysfs_led{SysfsLOC: "/beaglebone:green:usr2", Name: "usr2"}}
	Usr3 = LED{Sys: sysfs_led{SysfsLOC: "/beaglebone:green:usr3", Name: "usr3"}}
)

//GetLEDSysfs - returns complete abs path to gpio dir
//
func (l *LED) GetLEDSysfs() string {
	return ledBase + l.Sys.SysfsLOC
}

//LedOff - turns LED off
//
func (l *LED) LedOff() {
	fmt.Printf("LED %s off\n", l)
	WriteAndClose(l.GetLEDSysfs()+"/brightness", "0")

}

//LedOn - turns LED on
//
func (l *LED) LedOn() {
	fmt.Printf("LED %s on\n", l)
	WriteAndClose(l.GetLEDSysfs()+"/brightness", "1")
}
