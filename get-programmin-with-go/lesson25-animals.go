package main

import (
	"fmt"
	"math/rand"
	"time"
)

// code for monkey
type monkey struct {
	name string
}

func (m monkey) String() string {
	return m.name
}
func (m monkey) move() string {
	switch rand.Intn(2) {
	case 0:
		return "jumps"
	case 1:
		return "swings"
	default:
		return "!!!"
	}
}
func (m monkey) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "banana"
	case 1:
		return "orange"
	case 2:
		return "ant"
	default:
		return "!!!"
	}
}

// code for horse
type horse struct {
	name string
}

func (h horse) String() string {
	return h.name
}
func (h horse) move() string {
	switch rand.Intn(2) {
	case 0:
		return "jumps"
	case 1:
		return "runs"
	default:
		return "!!!"
	}
}
func (h horse) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "grass"
	case 1:
		return "apple"
	case 2:
		return "carrot"
	default:
		return "!!!"
	}
}

type animal interface {
	move() string
	eat() string
}

func doSomething(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%s %s \n", a, a.move())
	case 1:
		fmt.Printf("%s eats %s \n", a, a.eat())
	}
}

const (
	sunset, sunrise int = 18, 6
)

func main() {
	animals := []animal{
		monkey{"jeffry"},
		horse{"james"},
	}
	var hour int

	// Simulate a 3-day cycle
	for day := 1; day <= 3; day++ {
		fmt.Printf("Day %d\n", day)
		hour = 0
		// Simulate 24 hours in a day
		// Animals are active between sunrise and sunset
		for i := 0; i < 24; i++ {
			hour++
			fmt.Printf("%02d:00 ", hour)
			if hour < sunrise || hour >= sunset {
				fmt.Printf("Animals are Sleeping\n")
			} else {
				i := rand.Intn(len(animals))
				doSomething(animals[i])
			}
			time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds to simulate time passing
		}
	}
}
