package main

import (
	"fmt"
	"reflect"
)

// O(N^3)
func threeSum(nums []int) []interface{} {
	ret := make([]interface{}, len(nums))
	if len(nums) < 3 {
		return ret
	}

	// filter_nums := removeDuplicateElement(nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					// fmt.Println([]int{nums[i], nums[j], nums[k]})
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return removeDuplicateElement(ret)
}

func removeDuplicateElement(nums []interface{}) []interface{} {
	hashTable := make(map[interface{}]int)
	retTable := make(map[int]interface{})
	length := 0

	for i, v := range nums {
		if _, ok := hashTable[v]; ok {
			fmt.Println(v, " exists")
			continue
		}
		hashTable[v] = i
		retTable[length] = v
		length++
	}

	ret := make([]interface{}, length)
	for i, v := range retTable {
		ret[i] = v
	}
	return ret
}

func main() {
	test1 := []int{-1, 0, 1, 2, -1, -4}
	ret := make([]interface{}, 3)

	fmt.Println(reflect.TypeOf(test1))
	fmt.Println(reflect.TypeOf(ret))

	// fmt.Println(removeDuplicateElement(test1))
	// ret := threeSum(test1)
	// fmt.Println(ret)

	// for _, v := range ret {
	// 	fmt.Println(v)
	// }
}
