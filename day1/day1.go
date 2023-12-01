package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve() (int, int) {

	var solutionString string
	var secondSolutionString string

	var solution int
	var secondSolution int

	file, err := os.Open("day1/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// first half
		var solutionInt int
		solutionString = addFirstLast(numbersInString(scanner.Text()))
		solutionInt, err = strconv.Atoi(solutionString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		solution += solutionInt

		// second half
		// clean up the string first with cleanupInputRow()
		// this turns "one2three" -> "123"
		var secondSolutionInt int
		cleanString := cleanupInputRow(scanner.Text())

		secondSolutionString = addFirstLast(numbersInString(cleanString))
		secondSolutionInt, err = strconv.Atoi(secondSolutionString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		secondSolution += secondSolutionInt
	}
	return solution, secondSolution
}
