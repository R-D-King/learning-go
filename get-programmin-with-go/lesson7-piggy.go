package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello World")
	saveMoneyByCents()
}

func saveMoneyByCents() {
	var balance int64
	for balance < 2000 {
		fmt.Printf("Your balance Is: %v $\n", balance/100)

		switch deposit := rand.Intn(3); deposit {
		case 0:
			balance += 5
		case 1:
			balance += 10
		case 2:
			balance += 25
		}
	}
	fmt.Printf("Balance is: %v $\n", balance/100)
	fmt.Println("You Reached Your Goal!!!")
}
