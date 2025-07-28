package main

import "fmt"

func main() {
	sli := []string{}
	lastCap := cap(sli)

	for i := 0; i < 10000; i++ {
		sli = append(sli, "element")
		if cap(sli) != lastCap {
			fmt.Println(cap(sli))
			lastCap = cap(sli)
		}
	}
}
