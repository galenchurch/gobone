package gobone

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type GPIO struct {
	Sys   sysfs_gpio
	Dir   string //"in or "out"
	Value string //"0" or "1"
	PUPD  string //
}

//PinMap - GPIO Pins are passed in the form:
// "1_12" := bank 1 pin 12
//
func PinMap(n string) (g sysfs_gpio, err error) {
	s := regexp.MustCompile("_").Split(n, 2)

	g.Name = n

	if g.Bank, err = strconv.Atoi(s[0]); err != nil {
		log.Printf("Error Bank Conversion: %s", err)
		return g, err
	}

	if g.Number, err = strconv.Atoi(s[1]); err != nil {
		log.Printf("Error Number Conversion: %s", err)
		return g, err
	}

	//Calculate IO line
	dig := g.Bank*32 + g.Number
	g.SysfsLOC = "/gpio" + strconv.Itoa(dig)

	return g, nil
}

func InitGPIO(name string) (g GPIO, err error) {
	g.Sys, err = PinMap(name)
	return g, err
}

//SetDir - Sets GPIO Direction
//val = { "in", "out"}
func (g *GPIO) SetDir(val string) {
	g.Dir = "out"
	WriteAndClose(g.GetGPIOSysfs()+"/direction", val)
}

//SetVal - Sets GPIO value
//val = { "1", "0"}
//TODO: add varification of val
func (g *GPIO) SetVal(val string) {
	if g.Dir == "out" {
		WriteAndClose(g.GetGPIOSysfs()+"/value", val)
	} else if g.Dir == "in" {
		log.Printf("Error %s is configured an Input", g.Sys.Name)
	} else {
		log.Printf("Error %s is broke", g.Sys.Name)
	}
}

//ReadVal - Read GPIO value
//val = { "1", "0"}
//
func (g *GPIO) ReadVal() (r []byte) {
	r = ReadByteAndClose(g.GetGPIOSysfs() + "/value")
	fmt.Printf("Read: %s\n", r)
	return r
}

//GetGPIOSysfs - returns complete abs path to gpio dir
//
func (g *GPIO) GetGPIOSysfs() string {
	return gpioBase + g.Sys.SysfsLOC
}
