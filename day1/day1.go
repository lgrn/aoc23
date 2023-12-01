package day1

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() (int, int) {
	file, err := os.Open("day1/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(numbersInString(scanner.Text()))

		//if scanner.Text() != "" {
		//	// non-empty line detected, convert to int
		//	val, err := strconv.Atoi(scanner.Text())
		//
		//	if err != nil {
		//		fmt.Println(err)
		//		os.Exit(1)
		//	}
		//
		//	sum += val
		//} else if scanner.Text() == "" {
		//	// empty line detected, this means end of set
		//	// append 'sum' to the list of sums
		//	sums = append(sums, sum)
		//	sum = 0
		//}
	}

	// all lines have been processed

	//	threeBiggest := findNLargestNumbers(sums, 3)
	//
	//	return findLargestNumber(sums), sumSlice(threeBiggest)
	os.Exit(0)
}
