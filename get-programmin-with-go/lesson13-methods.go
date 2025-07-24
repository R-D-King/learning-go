package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64
type kelvin float64

func main() {
	var c celsius = 127
	var k kelvin = 127
	var f fahrenheit = 127

	fmt.Println(c.kelvin(), "Kelvin")
	fmt.Println(c.fahrenheit(), "Fahrenheit")

	fmt.Println(k.celsius(), "Celsius")
	fmt.Println(k.fahrenheit(), "Fahrenheit")

	fmt.Println(f.celsius(), "Celsius")
	fmt.Println(f.kelvin(), "Kelvin")
}

// from c to k
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// from c to f
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

// from k to c
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// from k to f
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(((k - 273.15) * 9.0 / 5.0) + 32)
}

// from f to c
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5.0 / 9.0)
}

// from f to k
func (f fahrenheit) kelvin() kelvin {
	return kelvin(((f - 273.15) * 5.0 / 9.0) + 273.15)
}
