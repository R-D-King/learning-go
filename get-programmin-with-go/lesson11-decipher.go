package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	decryptedText := "" // Initialize an empty string to build the result

	keywordLen := len(keyword) // Get the length of the keyword once
	for i, c := range cipherText {
		// i: current index (0, 1, 2, ...)
		// c: current rune (e.g., 'C', 'S', 'O', ...)

		// 1. Determine the corresponding keyword character
		// We use the modulo operator (%) to cycle through the keyword
		// If i is 0, 0 % 6 = 0 (keyword[0] -> 'G')
		// If i is 5, 5 % 6 = 5 (keyword[5] -> 'G')
		// If i is 6, 6 % 6 = 0 (keyword[0] -> 'G') - This is where it wraps around!
		k := rune(keyword[i%keywordLen])

		// 2. Convert characters to their numerical values (0-25)
		// 'A' is treated as 0, 'B' as 1, etc.
		cipherInt := charToInt(c) // e.g., 'C' - 'A' = 2
		keyInt := charToInt(k)    // e.g., 'G' - 'A' = 6

		// 3. Apply the Vigenere decryption formula
		// Pi = (Ci - Ki + 26) mod 26
		// The '+ 26' is crucial here. Why?
		// Let's say Ci = 2 ('C') and Ki = 6 ('G').
		// If we just do (2 - 6) = -4. A negative modulo result is problematic.
		// (-4 + 26) % 26 = 22 % 26 = 22. This corresponds to 'W'.
		// If Ci = 10 ('K') and Ki = 1 ('B').
		// (10 - 1 + 26) % 26 = 35 % 26 = 9. This corresponds to 'J'.
		// This ensures the result is always a positive integer between 0 and 25,
		// fulfilling the "no if statements" requirement.
		decryptedInt := (cipherInt - keyInt + 26) % 26

		// 4. Convert the numerical result back to an uppercase letter
		// e.g., 22 becomes 'A' + 22 = 'W'
		// Add the converted character to our result string
		decryptedText += string(intToChar(decryptedInt))
	}

	fmt.Println(decryptedText) // Print the final decrypted message
}

// Helper functions: These are essential for converting between character and integer representations.
// They make the main logic cleaner and more readable.

// charToInt converts an uppercase letter rune (e.g., 'A', 'B') to an integer (0-25).
func charToInt(r rune) int {
	return int(r - 'A') // Simple arithmetic on ASCII/Unicode values
}

// intToChar converts an integer (0-25) back to an uppercase letter rune.
func intToChar(i int) rune {
	return rune('A' + i) // Simple arithmetic to get the character
}

//package main
//
//import "fmt"
//
//func main() {
//	cipherText := "CSOITEUIWUIZNSROCNKFD"
//	keyword := "GOLANG"
//	decryptedText := make([]rune, len(cipherText)) // Use a rune slice for efficiency
//
//	keywordLen := len(keyword)
//	for i := 0; i < len(cipherText); i++ {
//		// Get the corresponding character from cipherText and keyword as byte
//		cipherByte := cipherText[i]
//		keyByte := keyword[i % keywordLen]
//
//		// Convert bytes to their integer equivalents (0-25)
//		cipherInt := int(cipherByte - 'A')
//		keyInt := int(keyByte - 'A')
//
//		// Decryption formula: P = (C - K + 26) mod 26
//		decryptedInt := (cipherInt - keyInt + 26) % 26
//
//		// Convert back to byte/rune and store
//		decryptedText[i] = rune('A' + decryptedInt)
//	}
//
//	fmt.Println(string(decryptedText)) // Convert the rune slice back to a string for printing
//}
//// Helper functions: These are essential for converting between character and integer representations.
//// They make the main logic cleaner and more readable.
//
//// charToInt converts an uppercase letter rune (e.g., 'A', 'B') to an integer (0-25).
//func charToInt(r rune) int {
//	return int(r - 'A') // Simple arithmetic on ASCII/Unicode values
//}
//
//// intToChar converts an integer (0-25) back to an uppercase letter rune.
//func intToChar(i int) rune {
//	return rune('A' + i) // Simple arithmetic to get the character
//}
