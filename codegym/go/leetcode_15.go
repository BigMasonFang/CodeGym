package main

import (
	"fmt"
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

// O(N^2) O(log(N))
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func main() {
	test2 := []int{-1, 0, 1, 2, -1, -4, 5, 9, 12, 1, 0, 68, 20}
	ret1 := threeSum(test2)
	fmt.Println(ret1)

}
