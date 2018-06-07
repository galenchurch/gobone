package gobone

const gpioBase string = "/sys/class/gpio"

type sysfs_gpio struct {
	Name     string
	Bank     int
	Number   int
	SysfsLOC string
}

const ledBase string = "/sys/class/leds"
