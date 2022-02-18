package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		} else {
			m[v] = i
		}
	}
	return false
}
