package main

import (
	"fmt"
)

func permutation(src []int) {
	for i := 0; i < len(src); i++ {
		first := src[i]
		temp := make([]int, 0)
		temp = append(temp, src[:i]...)
		temp = append(temp, src[i+1:]...)
		order([]int{first}, temp)
	}
}

func order(s []int, src []int) {
	fmt.Println("in the recur ", s, src)
	if len(src) <= 1 {
		fmt.Println("result: ", s, src, "\n")
		return
	}
	first := src[0]
	for i := 0; i < len(src); i++ {
		// swap
		fmt.Println(i)
		first, src[i] = src[i], first
		temp := make([]int, 0)
		temp = append(temp, src[1:]...)
		s_demo := make([]int, len(s))
		s_demo = s
		dd := make([]int, 0)
		dd = append(s_demo, first)
		fmt.Println("befor enter next recur ", dd, temp)
		order(dd, temp)
	}
}

func main() {
	permutation([]int{1, 2, 3, 4, 5})
}
