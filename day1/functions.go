package day1

import "sort"

// functions

// numbersInString() takes a string and returns a slice of ints being all
// numbers found in that string
func numbersInString() (s string) []int {
	var numbers []int
	for _, char := range s {
		if unicode.IsDigit(char) {
			numbers = append(numbers, int(char))
		}
	}
	return numbers
}
