package main

import "fmt"

// switch case
func checkSeason(s int) string {
	switch s {
	case 12, 1, 2:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	default:
		return "error input"
	}
}

// reverse array
func bubbleSort(slice []int) {
	// loop time n - 1
	for i := 0; i < len(slice)-1; i++ {
		// loop time n -i -1
		for j := len(slice) - 1; j > i; j-- {
			if slice[i] >= slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// test bubble sort
func testbubbleSort() {
	var testSlice1 = []int{7, 8, 9, 1, 2, 7}
	bubbleSort(testSlice1)
	fmt.Println(testSlice1)
}

func main() {
	fmt.Println(checkSeason(10))
	// fmt.Println(checkSeason(7))
	// fmt.Println(checkSeason(4))
	testbubbleSort()
}
