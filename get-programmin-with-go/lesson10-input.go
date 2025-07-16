package main

import "fmt"

func main() {
	var input string
	input = "Hell Yeah!!!!!!"

	var status bool
	switch input {
	case "0", "false", "no":
		status = false
		fmt.Println(status)
	case "1", "true", "yes":
		status = true
		fmt.Println(status)
	default:
		fmt.Println("Invalid input")
	}

}
