package main

import (
	"fmt"
	s "strings"
)

type Solution struct {
	a string
}

func (solution Solution) print() {
	fmt.Println("u are sexy")
}

func (solution Solution) repeatPrint(text string, x int) {
	// this is notation
	for i := 0; i < x; i++ {
		fmt.Println(s.Join([]string{solution.a, text}, " "))
	}
}

func main() {
	s := Solution{a: "fx"}
	text := "hello"

	// declare variables
	a, b := 2, 3
	fmt.Println(a, b)
	var e, f, g int = 4, 5, 6
	fmt.Println(e, f, g)

	s.print()
	s.repeatPrint(text, 3)
}
