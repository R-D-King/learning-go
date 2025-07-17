package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	processedPlainText := preparePlaintext(plainText)

	fmt.Println(encrypt(processedPlainText, keyword))
}

func preparePlaintext(text string) string {
	text = strings.Replace(text, " ", "", -1)
	text = strings.ToUpper(text)
	return text
}

func encrypt(PlainText, keyword string) string {
	cipherText := ""
	keywordLen := len(keyword)
	for i, p := range PlainText {
		pInt := int(p - 'A')
		kInt := int(keyword[i%keywordLen] - 'A')
		cInt := (pInt + kInt) % 26
		cipherText += string('A' + cInt)
	}
	return cipherText
}
