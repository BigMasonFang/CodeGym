package main

import (
	"fmt"
)

// O(N^2)
func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	for i := 1; i < len(nums); i++ {
		if (nums[i-1] + nums[i]) == target {
			ret[0], ret[1] = i-1, i
			return ret
		}
	}
	return nil
}

// O(N)
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	// test1 := []int{2, 7, 11, 5}
	test2 := []int{3, 2, 4}

	fmt.Println(twoSum(test2, 6))
}
