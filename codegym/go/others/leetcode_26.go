package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	// slow,fast
	s, f := 0, 1
	for f < l {
		if nums[s] != nums[f] {
			s++
			nums[s] = nums[f] // assignment
		}
		f++
	}
	return s + 1

	// n := len(nums)
	// if n == 0 {
	//     return 0
	// }
	// slow := 1
	// for fast := 1; fast < n; fast++ {
	//     if nums[fast] != nums[fast-1] {
	//         nums[slow] = nums[fast]
	//         slow++
	//     }
	// }
	// return slow
}

func main() {
	// test := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	test2 := []int{1, 2, 3, 4, 5}
	fmt.Println(removeDuplicates(test2))
	fmt.Println(test2)
}
