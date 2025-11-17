package main

import (
	"errors"
	"fmt"
)

const (
	empty         = 0
	rows, columns = 9, 9
)

type Cell struct {
	digit int
	fixed bool
}
type Grid [rows][columns]Cell

// Errors that could occur.
var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in this row")
	ErrInColumn   = errors.New("digit already present in this column")
	ErrInRegion   = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			if d != empty {
				grid[r][c].digit = int(d)
				grid[r][c].fixed = true
			}
		}
	}
	return &grid
}

// Set a digit on a Sudoku grid.
func (g *Grid) Set(row, column int, digit int8) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.isFixed(row, column):
		return ErrFixedDigit
	case g.inRow(row, digit):
		return ErrInRow
	case g.inColumn(column, digit):
		return ErrInColumn
	case g.inRegion(row, column, digit):
		return ErrInRegion
	}
	g[row][column].digit = int(digit)
	return nil
}

// Clear a cell from the Sudoku grid.
func (g *Grid) Clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.isFixed(row, column):
		return ErrFixedDigit
	}
	g[row][column].digit = empty
	return nil
}
func inBounds(row, column int) bool {
	return row >= 0 && row < rows && column >= 0 && column < columns
}
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}
func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == int(digit) {
			return true
		}
	}
	return false
}
func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == int(digit) {
			return true
		}
	}
	return false
}
func (g *Grid) inRegion(row, column int, digit int8) bool {
	startRow := (row / 3) * 3
	startCol := (column / 3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if g[r][c].digit == int(digit) {
				return true
			}
		}
	}
	return false
}
func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}
func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	err := s.Set(0, 5, 2)
	if err != nil {
		fmt.Println(err)
	}

	// for _, row := range s {
	// 	fmt.Println(row)
	// }

	// Loop through each row and then each cell to print just the digit
	fmt.Println("Current Sudoku Grid:")
	for _, row := range s {
		for _, cell := range row {
			fmt.Printf("%d ", cell.digit) // Print the digit and a space
		}
		fmt.Println() // Print a newline at the end of each row
	}
}
