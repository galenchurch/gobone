package gobone

import (
	"fmt"
	"log"
	"os"
)

const gpioBase string = "/sys/class/gpio"

type sysfs_gpio struct {
	Name     string
	Bank     int
	Number   int
	SysfsLOC string
}

const ledBase string = "/sys/class/leds"

type sysfs_led struct {
	Name     string
	SysfsLOC string
}

//WriteAndClose - Writes v to file passed as fi
//TODO add checks
func WriteAndClose(fi string, v string) {
	f, err := os.OpenFile(fi, os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(v)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Write to %s\n", fi)
}

//ReadByteAndClose - Returns []byte read from file fi
//
func ReadByteAndClose(fi string) []byte {

	var size int
	v := make([]byte, 1)
	f, err := os.OpenFile(fi, os.O_RDONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	size, err = f.Read(v)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %v from %s\nWas %d bytes in len", v, fi, size)
	return v
}
