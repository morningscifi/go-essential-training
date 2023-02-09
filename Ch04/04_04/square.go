package main

import (
	"fmt"
	"log"
)

// Square is a square
type Square struct {
	x      int
	y      int
	length int
}

// NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	if x <= 0 {
		return nil, fmt.Errorf("x must be positive")
	}
	if y <= 0 {
		return nil, fmt.Errorf("y must be positive")
	}
	s := Square{x, y, length}

	return &s, nil
}

// Move moves the square
func (s *Square) Move(dx int, dy int) {
	s.x += dx
	s.y += dy
}

// Area returns the square are
func (s *Square) Area() int {
	return s.length * s.length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
