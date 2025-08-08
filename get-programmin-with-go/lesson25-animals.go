package main

import (
	"fmt"
	"math/rand"
)

// code for monkey
type monkey struct {
	name string
}

func (m monkey) String() string {
	return m.name
}
func (m monkey) move() string {
	move := []string{"jumps", "swings"}
	switch rand.Intn(len(move)) {
	case 0:
		return move[0]
	case 1:
		return move[1]
	default:
		return "!!!"
	}
}
func (m monkey) eat() string {
	food := []string{"banana", "orange", "ant"}
	switch rand.Intn(len(food)) {
	case 0:
		return food[0]
	case 1:
		return food[1]
	case 2:
		return food[2]
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
	move := []string{"jumps", "runs"}
	switch rand.Intn(len(move)) {
	case 0:
		return move[0]
	case 1:
		return move[1]
	default:
		return "!!!"
	}
}
func (h horse) eat() string {
	food := []string{"grass", "apple", "carrot"}
	switch rand.Intn(len(food)) {
	case 0:
		return food[0]
	case 1:
		return food[1]
	case 2:
		return food[2]
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

	for day := 1; day <= 3; day++ {
		fmt.Printf("Day %d\n", day)
		hour = 0
		for i := 0; i < 24; i++ {
			hour++
			fmt.Printf("%02d:00 ", hour)
			if hour < sunrise || hour >= sunset {
				fmt.Printf("Animals are Sleeping\n")
			} else {
				i := rand.Intn(len(animals))
				doSomething(animals[i])
			}
		}
	}
}
