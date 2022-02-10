package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
func keyboardIo() {
	// use reflect to check type
	// os.Stdin utilize Read() func
	fmt.Println(reflect.TypeOf(os.Stdin))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	fmt.Print("You live in " + city)
}

func main() {
	printMultiplicationTable(10)
	keyboardIo()
}
