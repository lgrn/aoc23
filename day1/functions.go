package day1

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// functions

// numbersInString() takes a string and returns a slice of ints being all
// numbers found in that string
func numbersInString(s string) []int {
	var numbers []int
	for _, char := range s {
		if unicode.IsDigit(char) {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			numbers = append(numbers, digit)
		}
	}
	return numbers
}

// addFirstLast() takes the first and last integer of an arbitrarily long []int
// and *adds* the numbers together (not sum), i.e. 1+2=12 and not 3.
func addFirstLast(numbers []int) (sum string) {
	// switch on the length of numbers: it's either one, two or more.
	switch len(numbers) {
	case 0:
		os.Exit(1)
	case 1:
		return strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[0])
	case 2:
		return strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1])
	}
	return strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1])
}

// cleanupInputRow() takes an input row, but looks for numbers expressed
// alphabetically, so that "xxone2three" returns "123"
func cleanupInputRow(row string) string {
	numberNameMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var resultstring string
	stringLength := len(row)

	for index, char := range row {
		// we are looking at one character in the string.
		// does this already look like an integer? if so, just append it
		if _, err := strconv.Atoi(string(char)); err == nil {
			resultstring += string(char)
		} else {
			// look forward 3,4 and 5 characters and see if that matches
			// something in our translation map.
			for i := 2; i <= 6; i++ {
				// do not overshoot the length of the string
				if index+i <= stringLength {
					for key, _ := range numberNameMap {
						if key == row[index:index+i] {
							resultstring += numberNameMap[key]
							break
						}
					}
				}
			}
		}
	}
	return resultstring
}
