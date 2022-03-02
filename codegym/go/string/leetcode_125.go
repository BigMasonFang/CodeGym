package main

import (
	"strings"
	"unicode"
)

// old way 100/25.6
func isPalindrome(s string) bool {
	// preprocess input
	// var str_build strings.Builder
	var process []rune

	for _, c := range s {
		if unicode.IsLetter(c) {
			process = append(process, unicode.ToLower(c))
		} else if unicode.IsNumber(c) {
			process = append(process, c)
		}
	}
	// processStr := str_build.String()
	// processStr := string(process)

	i, j, compareTimes := 0, len(process)-1, len(process)/2
	for compareTimes > 0 {
		if process[i] != process[j] {
			return false
		}
		compareTimes--
		i++
		j--
	}
	return true
}

// standard
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
