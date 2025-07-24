package main

import "fmt"

func main() {
	var board [8][8]rune
	board[0] = [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}
	board[7] = [8]rune{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}
	drawBoard(board)
}

func drawBoard(board [8][8]rune) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if board[x][y] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c ", board[x][y])
			}
		}
		fmt.Printf("\n")
	}
}
