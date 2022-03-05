package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	// greedy algo
	left, right, flagLeft, flagRight := 0, len(s)-1, true, true
	for left < right {
		if s[left] != s[right] {
			// do not repeatly use left, right in outside loop
			// because u need to loop twice
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flagLeft = false
					break
				}
			}
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flagRight = false
					break
				}
			}
			return flagLeft || flagRight
		} else {
			left++
			right--
		}
	}
	return true
}

func main() {
	testSrt := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	rst := validPalindrome(testSrt)
	fmt.Println(rst)
}
