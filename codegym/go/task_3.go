package main

import (
	"fmt"
)

// multiplication table
func printMultiplicationTable(num int) {
	line := make([]int, num)
	for i := 0; i < num; i++ {
		line[i] = i + 1
	}
	for _, n := range line {
		for i := 0; i < num; i++ {
			fmt.Print(n*(i+1), " ")
		}
		fmt.Println()
	}
}

// keyborad io
func io() {

}

func main() {
	printMultiplicationTable(10)
}
