package main

import (
	"fmt"
)

type item struct {
	name string
}
type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up %v\n", c.name, i.name)
	c.leftHand = i
}
func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v's left hand is full\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives %v to %v\n", c.name, to.leftHand.name, to.name)
}
func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is empty-handed", c.name)
	}
	return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}
func main() {
	michael := &character{name: "Michael"}
	apple := &item{name: "apple"}
	michael.pickup(apple)

	kinght := &character{name: "Knight"}
	michael.give(kinght)

	fmt.Println(michael)
	fmt.Println(kinght)
}
