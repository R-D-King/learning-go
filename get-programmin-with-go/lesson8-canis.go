package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792) // km/s
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("236000000000000000", 0)
	fmt.Println("Canis Major Dwarf is", distance, "km away.")

	// Calculate the time it would take to travel to Canis Major Dwarf at light speed
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	// Convert seconds to light years
	secondsPerYear := new(big.Int).Mul(secondsPerDay, big.NewInt(365))
	lightYears := new(big.Int)
	lightYears.Div(seconds, secondsPerYear)

	fmt.Println("That is", lightYears, "light years away.")
}
