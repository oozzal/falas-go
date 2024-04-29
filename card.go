package main

import "fmt"

var (
	faces  = [4]string{"♠️", "♣️", "❤️️", "♦️"}
	values = [13]Value{
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
		{"J", 11},
		{"Q", 12},
		{"K", 13},
		{"A", 14},
	}
)

type Value struct {
	name string
	rank int
}

type Card struct {
	face  string
	value Value
}

func (c Card) Rank() int {
	return c.value.rank
}

func (c Card) Display() string {
	return fmt.Sprintf("%s %s", c.face, c.value.name)
}
