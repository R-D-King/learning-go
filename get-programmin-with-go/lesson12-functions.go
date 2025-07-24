package main

import (
	"fmt"
)

func main() {
	input := 233.0
	fmt.Println("input is in Kelvin:", input)
	fmt.Println("in Celsius =", kelvinToCelsius(input))
	fmt.Println("in Fahrenheit =", kelvinToFahrenheit(input))
}

// kelvinToCelsius takes a temperature in Kelvins and converts it Celsius
func kelvinToCelsius(kelvin float64) float64 {
	kelvin -= 273.15
	return kelvin
}

// celsiusToFahrenheit takes temperature in Celsius and converts it Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	celsius = (celsius * 9.0 / 5.0) + 32.0
	return celsius
}

// kelvinToFahrenheit takes temperature in Kelvin and converts it to Celsius and then to Fahrenheit
func kelvinToFahrenheit(kelvin float64) float64 {
	kelvin = celsiusToFahrenheit(kelvinToCelsius(kelvin))
	return kelvin
}
