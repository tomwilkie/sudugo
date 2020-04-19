package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	errOutOfRange      = errors.New("out of range")
	errNewlineExpected = errors.New("expected newline")
)

const (
	topRow    = "┌───────┬───────┬───────┐\n"
	midRow    = "├───────┼───────┼───────┤\n"
	bottomRow = "└───────┴───────┴───────┘\n"
	rowStart  = "│ "
	rowEnd    = "│\n"
)

type suduko [9][9]int

func (s *suduko) String() string {
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

func (s *suduko) read(r io.Reader) error {
	br := bufio.NewReader(r)

	// Read: "┌───────┬───────┬───────┐\n"
	if s, err := br.ReadString('\n'); err != nil {
		return err
	} else if s != topRow {
		return fmt.Errorf("unexpected row: %s", s)
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
