package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	// O(N^2)
	// length := 0
	// for i := len(nums) - 1; i > 0; i-- {
	// 	for j := i - 1; j >= 0; j-- {
	// 		if nums[i] == nums[j] {
	// 			nums = append(nums[:j], nums[i:]...)
	// 			length++
	// 		}
	// 	}
	// }
	// return length

	left, right := 0, 1

	for right < len(nums)-1 {
		if nums[left] == nums[right] {
			// compare and reassgin
			nums[right] = nums[right+1]
			right++
		} else {
			left++
			nums[left] = nums[right]
		}
	}
	// fmt.Println(left, right)
	if nums[left] != nums[right] {
		left++
		nums[left] = nums[right]
	}

	return left
}

func main() {
	test := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(test))
	fmt.Println(test)
}
