// three sum
package main

import (
	"fmt"
	"math"
	"sort"
)

// O(N^3)
func threeSum1(nums []int) [][]int {
	ret := make([][]int, len(nums))
	if len(nums) < 3 {
		return ret
	}
	ret_num := 0

	// filter_nums := removeDuplicateElement(nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					// fmt.Println([]int{nums[i], nums[j], nums[k]})
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
					ret_num++
				}
			}
		}
	}
	return removeDuplicateElement(ret[(len(ret) - ret_num):])
}

// O(N^2)
func removeDuplicateElement(nums [][]int) [][]int {
	for i := len(nums) - 1; i >= 0; i-- {
		j := i - 1
		for j >= 0 {
			if compareElement(nums[i], nums[j]) {
				nums = append(nums[:j], nums[j+1:]...)
			}
			j--
		}
	}
	return nums
}

func compareElement(e1, e2 []int) bool {
	// order does not matter
	hashTable := map[int]int{}
	for i, v := range e1 {
		hashTable[v] = i
	}
	for _, v := range e2 {
		if _, ok := hashTable[v]; !ok {
			return false
		}
	}
	return true
}

// O(N^2)+O(log(N))
// 20220220 revisit
// 这道题主要在于 1排序之后使用双指针 2 重复元素的处理
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n <= 2 {
		return [][]int{}
	} else if n == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{[]int{nums[0], nums[1], nums[2]}}
		}
	}
	sort.Ints(nums)
	ret := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		// nums handle repeat value in the past
		// this repeat is "element counted"
		// and
		// nums[i] > 0 and sorted, then must >0
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		left, right := i+1, n-1

		for left < right {
			// three conditions
			if nums[i]+nums[left]+nums[right] == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				// deal with repeat elements left side
				for left < right && nums[left] == nums[left+1] {
					left++
					// out of boundary
					if left == n-2 {
						break
					}
				}
				// deal with repeat elments right side
				for left < right && nums[right] == nums[right-1] {
					right--
					if right == n-2 {
						break
					}
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ret
}

func main() {
	// test2 := []int{-1, 0, 1, 2, -1, -4, 5, 9, 12, 1, 0, 68, 20}
	// ret1 := threeSum(test2)

	fmt.Println(math.Abs(-33))
	// fmt.Println(ret1)

}
