package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// header
	header := fmt.Sprintf("%-20v %5v %-15v %10v", "SpaceLine", "Days", "Trip Type", "Price ($M)")
	fmt.Println(header)
	separator := strings.Repeat("=", len(header))
	fmt.Println(separator)

	for i := 10; i > 0; i-- {
		ticketGenerator()
	}

}

func ticketGenerator() {

	// // departure date
	// year := 2020
	// month := 10
	// day := 13
	hoursPerDay := 24
	distance := 62100000        // km
	speed := rand.Intn(14) + 16 // 16-30 km/s

	price := 36 + (speed - 16)

	// calculate the duration
	duration := distance / (speed * (60 * 60 * hoursPerDay)) //in days
	var spacelines string
	var tripType string

	switch companyID := rand.Intn(3); companyID {
	case 0:
		spacelines = "Space Adventures"
	case 1:
		spacelines = "SpaceX"
	case 2:
		spacelines = "Virgin Galactic"
	}

	switch tripCode := rand.Intn(2); tripCode {
	case 0:
		tripType = "one-way"
	case 1:
		tripType = "Round-trip"
		price *= 2
	}

	fmt.Printf("%-20v %5v %v %-15v $%4v\n", spacelines, duration, "days", tripType, price)
}
