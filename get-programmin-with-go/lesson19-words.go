package main

import (
	"fmt"
	"strings"
)

const punctuationCutset = `!"#$%&'()*+,-./:;<=>?@[\]^_` + "`{|}~"

func main() {
	sampleText := "As far as eye could reach he saw nothing but the stems of the great plants about him\nreceding in the violet shade, and far overhead the multiple transparency of huge leaves\nfiltering the sunshine to the solemn splendour of twilight in which he walked. Whenever\nhe felt able he ran again; the ground continued soft and springy, covered with the same\nresilient weed which was the first thing his hands had touched in Malacandra. Once or\ntwice a small red creature scuttled across his path, but otherwise there seemed to be no\nlife stirring in the wood; nothing to fear—except the fact of wandering unprovisioned\nand alone in a forest of unknown vegetation thousands or millions of miles beyond the\nreach or knowledge of man."
	frequencies := countWordFrequency(sampleText)
	for word, count := range frequencies {
		fmt.Printf("'%s': %d\n", word, count)
	}
}

func countWordFrequency(text string) map[string]int {
	wordCount := make(map[string]int)
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	for _, word := range words {
		trimmedWord := strings.Trim(word, punctuationCutset)
		if trimmedWord != "" {
			wordCount[trimmedWord]++
		}
	}
	return wordCount
}
