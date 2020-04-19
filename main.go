package main

import (
	"fmt"
	"log"
	"os"
)

func solve(s *suduko) error {
	return nil
}

func main() {
	var s suduko

	if err := s.read(os.Stdin); err != nil {
		log.Fatal("Failed to read suduko from stdin, error:", err)
	}

	if err := solve(&s); err != nil {
		log.Fatal("Failed to solve suduko, error:", err)
	}

	fmt.Println(s.String())
}
