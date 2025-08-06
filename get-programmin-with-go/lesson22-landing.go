package main

import "fmt"

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	lat, long float64
}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	Spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	Opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	Curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	InSight := newLocation(coordinate{4, 30, 0, 'N'}, coordinate{135, 54, 0, 'E'})
	fmt.Printf("Spirit: %.6f\n", Spirit)
	fmt.Printf("Opportunity: %.6f\n", Opportunity)
	fmt.Printf("Curiosity: %.6f\n", Curiosity)
	fmt.Printf("InSight: %.6f\n", InSight)
}
