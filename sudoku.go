package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	topRow    = "┌───────┬───────┬───────┐\n"
	midRow    = "├───────┼───────┼───────┤\n"
	bottomRow = "└───────┴───────┴───────┘\n"
	rowStart  = "│ "
	rowEnd    = "│\n"
)

var errOutOfRange = errors.New("out of range")

// soduko is indexed by row x column.
type sudoku [9][9]int

func (s *sudoku) String() string {
	var b strings.Builder

	b.WriteString(topRow)
	for i := 0; i < 9; i++ {
		b.WriteString(rowStart)
		for j := 0; j < 9; j++ {
			if s[i][j] == 0 {
				b.WriteRune('.')
			} else {
				b.WriteString(strconv.Itoa(s[i][j]))
			}

			b.WriteRune(' ')
			if j%3 == 2 && j != 8 {
				b.WriteString(rowStart)
			}
		}
		b.WriteString(rowEnd)

		if i%3 == 2 && i != 8 {
			b.WriteString(midRow)
		}
	}
	b.WriteString(bottomRow)

	return b.String()
}

func (s *sudoku) read(r io.Reader) error {
	br := bufio.NewReader(r)

	for {
		// Read: "┌───────┬───────┬───────┐\n"
		if s, err := br.ReadString('\n'); err != nil {
			return err
		} else if strings.HasPrefix(s, "#") {
			continue
		} else if s == topRow {
			break
		} else {
			return fmt.Errorf("unexpected row: %s", s)
		}
	}

	for i := 0; i < 9; i++ {

		// Read "│ "
		if s, err := br.ReadString(' '); err != nil {
			return err
		} else if s != rowStart {
			return fmt.Errorf("unexpected row start: %s", s)
		}

		for j := 0; j < 9; j++ {
			c, err := br.ReadString(' ')
			if err != nil {
				return err
			}

			if len(c) != 2 {
				return errOutOfRange

			} else if c == ". " {
				(*s)[i][j] = 0

			} else {
				n, err := strconv.Atoi(c[:1])
				if err != nil {
					return err
				}

				if n < 0 || n > 9 {
					return errOutOfRange
				}

				s[i][j] = n
			}

			// Read "│ "
			if j%3 == 2 && j != 8 {
				if s, err := br.ReadString(' '); err != nil {
					return err
				} else if s != rowStart {
					return fmt.Errorf("unexpected row mid: %s", s)
				}
			}
		}

		// Read "|\n"
		if s, err := br.ReadString('\n'); err != nil {
			return err
		} else if s != rowEnd {
			return fmt.Errorf("unexpected row end: %#v", s)
		}

		// Read: "├───────┼───────┼───────┤\n"
		if i%3 == 2 && i != 8 {
			if s, err := br.ReadString('\n'); err != nil {
				return err
			} else if s != midRow {
				return fmt.Errorf("unexpected row: %s", s)
			}
		}
	}

	// Read: "└───────┴───────┴───────┘\n"
	if s, err := br.ReadString('\n'); err != nil {
		return err
	} else if s != bottomRow {
		return fmt.Errorf("unexpected row: %s", s)
	}

	return nil
}

func (s *sudoku) isPartialValid() bool {
	// Check bounds.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] < 0 || s[i][j] > 9 {
				return false
			}
		}
	}

	// Check rows.
	for i := 0; i < 9; i++ {
		var uniq [10]int
		for j := 0; j < 9; j++ {
			uniq[s[i][j]]++
		}
		for j := 1; j < 10; j++ {
			if uniq[j] > 1 {
				return false
			}
		}
	}

	// Check columns.
	for j := 0; j < 9; j++ {
		var uniq [10]int
		for i := 0; i < 9; i++ {
			uniq[s[i][j]]++
		}
		for i := 1; i < 10; i++ {
			if uniq[i] > 1 {
				return false
			}
		}
	}

	// Check squares.
	for i := 0; i < 9; i++ {
		var uniq [10]int
		for j := 0; j < 9; j++ {
			row := ((i / 3) * 3) + (j / 3)
			col := ((i % 3) * 3) + (j % 3)
			uniq[s[row][col]]++
		}
		for j := 1; j < 10; j++ {
			if uniq[j] > 1 {
				return false
			}
		}
	}

	return true
}

func (s *sudoku) isValid() bool {
	// Check it has no zeros.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == 0 {
				return false
			}
		}
	}

	return s.isPartialValid()
}
