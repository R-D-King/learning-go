package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y--
}
func (t *turtle) down() {
	t.x++
}
func (t *turtle) left() {
	t.x--
}
func (t *turtle) right() {
	t.x++
}

func main() {
	donatelo := turtle{}
	donatelo.up()
	donatelo.down()
	donatelo.left()
	fmt.Println(donatelo)
	donatelo.left()
	donatelo.left()
	donatelo.down()
	fmt.Println(donatelo)
	donatelo.down()
	donatelo.up()
	donatelo.right()
	fmt.Println(donatelo)
}
