package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

func main() {
	nums1 := []int{1, 2, 5, 0, 0, 0, 0}
	nums2 := []int{1, 1, 1, 4}
	// nums3 := []int{}

	merge(nums1, 3, nums2, 4)
	// merge(nums3, 0, nums3, 0)

	fmt.Println(nums1)
}
