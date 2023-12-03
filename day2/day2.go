package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() (int, int) {

	file, err := os.Open("day2/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var possibleGames []int
	var possibleGamesSum int
	var powerOfGames int

	for scanner.Scan() {
		if scanner.Text() != "" {
			// parse this row into a 'game' type
			game := parseRow(scanner.Text())
			// add possible game ids to 'possibleGames'
			if validateGame(game) {
				possibleGames = append(possibleGames, game.id)
			}
			// always add the power of this game to the sum of powers
			powerOfGames += powerOfCubeCount(getMaxCubeCount(game))
		}
	}

	for _, val := range possibleGames {
		possibleGamesSum += val
	}

	return possibleGamesSum, powerOfGames
}
