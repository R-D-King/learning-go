package main

import "strings"

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go splitWords(c0, c1)
	printWords(c1)

}
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}
func splitWords(upstream chan string, downstream chan string) {
	for v := range upstream {
		for _, words := range strings.Fields(v) {
			downstream <- words
		}
	}
	close(downstream)
}
func printWords(upstream chan string) {
	for v := range upstream {
		println(v)
	}
}
