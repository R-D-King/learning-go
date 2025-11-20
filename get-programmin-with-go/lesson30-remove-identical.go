package main

import (
	"fmt"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go source(c0)
	go filter(c0, c1)
	printingGopher(c1)

}
func source(c chan string) {
	for _, v := range []string{"a", "b", "b", "c", "d", "d", "d", "e"} {
		c <- v
	}
	close(c)
}
func filter(up, down chan string) {
	prev := ""
	for v := range up {
		if v != prev {
			down <- v
			prev = v
		}
	}
	close(down)
}
func printingGopher(c chan string) {
	for v := range c {
		fmt.Println(v)
	}
}
