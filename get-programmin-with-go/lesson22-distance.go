package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}
type location struct {
	name      string
	lat, long float64
}
type coordinate struct {
	d, m, s float64
	h       rune
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
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

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(name string, lat, long coordinate) location {
	return location{name, lat.decimal(), long.decimal()}
}

func main() {
	// The volumetric mean radius of various planets
	var earth = world{6371.0}
	var mars = world{3389.5}

	// Landing sites on mars
	Spirit := newLocation("Spirit", coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	Opportunity := newLocation("Opportunity", coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	Curiosity := newLocation("Curiosity", coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	InSight := newLocation("InSight", coordinate{4, 30, 0, 'N'}, coordinate{135, 54, 0, 'E'})

	landingSites := []location{Spirit, Opportunity, Curiosity, InSight}

	minDistance := math.MaxFloat64
	maxDistance := 0.0
	var closestLoc1, closestLoc2 location
	var farthestLoc1, farthestLoc2 location

	fmt.Println("Calculating distances between Mars landing sites (in km):")
	fmt.Println("-------------------------------------------------------")

	for i := 0; i < len(landingSites); i++ {
		// Start j from i+1 to avoid duplicate pairs and self-comparison
		for j := i + 1; j < len(landingSites); j++ {

			loc1 := landingSites[i]
			loc2 := landingSites[j]
			dist := mars.distance(loc1, loc2)
			fmt.Printf("Distance between %s and %s is %.2f km \n", loc1.name, loc2.name, dist)

			// Update the closest locations
			if dist < minDistance {
				minDistance = dist
				closestLoc1 = loc1
				closestLoc2 = loc2
			}

			// Update the farthest locations
			if dist > maxDistance {
				maxDistance = dist
				farthestLoc1 = loc1
				farthestLoc2 = loc2
			}
		}
	}

	fmt.Println("\n--- Summary ---")
	fmt.Printf("Closest locations: %s and %s (Distance: %.2f km)\n", closestLoc1.name, closestLoc2.name, minDistance)
	fmt.Printf("Farthest locations: %s and %s (Distance: %.2f km)\n", farthestLoc1.name, farthestLoc2.name, maxDistance)

	// Moving to Earth
	london := newLocation("London", coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := newLocation("Paris", coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	cairo := newLocation("Cairo", coordinate{30, 2, 0, 'N'}, coordinate{31, 14, 0, 'E'})
	myTown := newLocation("MyTown", coordinate{30, 6, 38.3, 'N'}, coordinate{31, 10, 17.3, 'E'})

	// Two more locations on mars
	mountSharp := newLocation("Mount Sharp", coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMons := newLocation("Olympus Mons", coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})

	// On Earth
	distLP := earth.distance(london, paris)
	distCM := earth.distance(cairo, myTown)

	// On Mars
	distMO := mars.distance(mountSharp, olympusMons)

	fmt.Printf("\n=======================================================\n")
	fmt.Printf("Distance between %s and %s is %.2f km \n", london.name, paris.name, distLP)
	fmt.Printf("Distance between %s and %s is %.2f km \n", cairo.name, myTown.name, distCM)
	fmt.Printf("Distance between %s and %s is %.2f km \n", mountSharp.name, olympusMons.name, distMO)
}
