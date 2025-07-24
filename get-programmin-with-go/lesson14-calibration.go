package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	var offset kelvin = 67
	sensor := calibrate(fakeSensor, offset)
	fmt.Println(sensor())
	fmt.Println(sensor())

}
