package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello World")
	saveMoney()
}

func saveMoney() {
	var balance float64
	for balance < 20 {
		fmt.Printf("Your balance Is: %.2f $\n", balance)

		switch deposit := rand.Intn(3); deposit {
		case 0:
			balance += 0.05
		case 1:
			balance += 0.10
		case 2:
			balance += 0.25
		}
	}
	fmt.Printf("Balance is: %.2f $\n", balance)
	fmt.Println("You Reached Your Goal!!!")
}
