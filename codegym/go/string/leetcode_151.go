package main

import (
	"fmt"
)

// own first try 31/24
func reverseWords(s string) string {
	wordsSlice := []string{}
	result := ""

	i, j := 0, 0
	for k := 0; k < len(s); k++ {
		// aim to pass the prefix space
		for s[k] == ' ' {
			if k == len(s)-1 {
				break
			}
			k++
		}
		i, j = k, k

		// aim to catch the word
		for s[j] != ' ' {
			if j == len(s)-1 {
				break
			}
			j++
		}
		// if i ==j means , in the last, no word available or a single char
		if i == j {
			if s[j] != ' ' {
				wordsSlice = append(wordsSlice, string(s[j]))
			}
			// fmt.Println(string(s[j]))
			break
		}
		// fmt.Println(wordsSlice)
		// fmt.Println(i, j, k)
		// here, when encounter the last word, slice index should add 1
		if j != len(s)-1 {
			wordsSlice = append(wordsSlice, string(s[i:j]))
		} else {
			if s[j] == ' ' {
				wordsSlice = append(wordsSlice, string(s[i:j]))
			} else {
				wordsSlice = append(wordsSlice, string(s[i:j+1]))
			}
			break
		}
		// here, j is set back to the white space in order to get the right k
		k = j
	}

	// fmt.Println(wordsSlice)
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
	// testStr2 := "  Bob    Love    Alice "
	// testStr3 := "  hello world      "
	testStr4 := " asdasd df f"
	result := reverseWords(testStr4)
	fmt.Println(result)
}
