package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func main() {
	a := newUniverse()
	b := newUniverse()

	// Seed the random number generator ONCE at the start of the program
	a.Seed()

	generations := 50
	for gen := 0; gen < generations; gen++ {
		// Adding a title
		genStr := fmt.Sprintf("GENERATION %d", gen)
		padding := (width - len(genStr)) / 2
		// Ensure we don't have negative padding or odd distribution
		leftPadding := strings.Repeat("-", padding)
		rightPadding := strings.Repeat("-", width-len(genStr)-padding)

		fmt.Printf("%s%s%s\n", leftPadding, genStr, rightPadding)

		// Showing the universe
		a.Show()
		time.Sleep(time.Millisecond * 1000)

		Step(a, b)

		a, b = b, a
	}

}

// Universe : Creates a new universe with the specified height and width
func newUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Show : Prints the universe to the screen, the alive cells are *, and the dead cells are " ".
func (u Universe) Show() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u[y][x] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Seed : Generates a randomly populated universe
func (u Universe) Seed() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if rand.Intn(4) == 0 {
				u[y][x] = true
			}
		}
	}
}

// Alive : Checks if the is alive
func (u Universe) Alive(x, y int) bool {
	x = (x%width + width) % width

	y = (y%height + height) % height

	return u[y][x]
}

// Neighbors : Calculating the live neighbors
func (u Universe) Neighbors(x, y int) int {
	liveNeighbors := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			// Skip the central cell itself
			if i == 0 && j == 0 {
				continue
			}

			neighborX := x + j
			neighborY := y + i

			if u.Alive(neighborY, neighborX) {
				liveNeighbors++
			}
		}
	}

	return liveNeighbors
}

// Next : Decides if the cell will be alive in the next generation
func (u Universe) Next(x, y int) bool {
	isAlive := u.Alive(x, y)
	liveNeighbors := u.Neighbors(x, y)

	if isAlive {
		if liveNeighbors < 2 || liveNeighbors > 3 {
			return false
		}
		return true
	} else {
		if liveNeighbors == 3 {
			return true
		}
		return false
	}

}

// Step : Reads through universe A while setting cells in universe B
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b[y][x] = a.Next(x, y)
		}
	}
}
