package main

import (
	"fmt"
)

func reverseWords(s string) string {
	wordsSlice := []string{}
	result := ""

	i, j := 0, 0
	for k := 0; k < len(s); k++ {
		if s[k] == ' ' {
			k++
			continue
		}
		i, j = k, k
		for s[j] != ' ' {
			j++
			if j == len(s)-1 {
				break
			}
		}
		fmt.Println(wordsSlice)
		// here, when encounter the last word, slice index should add 1
		if j != len(s)-1 {
			wordsSlice = append(wordsSlice, string(s[i:j]))
		} else {
			wordsSlice = append(wordsSlice, string(s[i:j+1]))
			break
		}
		// here, j is set back to the white space in order to get the right k
		k = j
	}

	for l := len(wordsSlice); l > 0; l-- {
		result += wordsSlice[l-1]
		if l != 1 {
			result += " "
		}
	}
	return result
}

func main() {
	// testStr := "the sky is blue"
	testStr2 := "  Bob      Love    Alice "
	result := reverseWords(testStr2)
	fmt.Println(result)
}
