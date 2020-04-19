package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var s sudoku

	if err := s.read(os.Stdin); err != nil {
		log.Fatal("Failed to read suduko from stdin, error:", err)
	}

	if !s.isPartialValid() {
		log.Fatal("Input suduko is invalid.")
	}

	if err := solve(&s); err != nil {
		log.Fatal("Failed to solve suduko, error:", err)
	}

	if !s.isValid() {
		log.Fatal("Failed to produce valid solution suduko")
	}

	fmt.Println("Solved:")
	fmt.Println(s.String())
}

func solve(s *sudoku) error {
	// First copy over all the fixed cells.
	result := *s

	for k := 0; k < 9*9; {
		i := k / 9
		j := k % 9

		// Skip solving if hit a "fixed" cell.
		if s[i][j] != 0 {
			k++
			continue
		}

		// Increment current cell; cell starts at 0, and it reset to 0 when
		// backtracking.
		result[i][j]++

		// Check result it valid - if so, move on.
		if result.isPartialValid() {
			k++
			continue
		}

		// If we still have some room to grow this cell, repeat.
		if result[i][j] < 9 {
			continue
		}

		// Need to backtrack.
		result[i][j] = 0
		for k > 0 {
			k--
			if s[k/9][k%9] == 0 {
				break
			}
		}
	}

	*s = result
	return nil
}
