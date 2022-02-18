package main

import (
	"fmt"
)

// O(N)
// delete (or append new slice) in reverse way
func removeElement(nums []int, val int) int {
	ret := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			ret--
		}
	}
	return ret
}

// O(N)
// double pointer
func removeElement2(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// improved O(N)
func removeElement3(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

func main() {
	test1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// ret := removeElement(test1, 2)
	ret := removeElement2(test1, 2)
	fmt.Println(ret, test1[:ret])
	// go slice index out of bound will not error, but return nil
	// fmt.Println(test1[4:])
}
