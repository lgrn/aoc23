package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A 'game' is a data structure representing one parsed line of a game and its
// reveals. It has an identifier (id: int) as well as a slice of maps of string
// -> int relationships, being color and amount.
//    Example:
//    "Game 69: 7 green, 4 blue, 3 red; 4 blue(...)"
//          ^^-- id
//              ^^^^^^^^^^^^^^^^^^^^^^-- reveal (one)
type game struct {
	id      int
	reveals []map[string]int
}

// parseRow parses a row of input and returns a 'game' struct with the game id
// and all reveals presented in that game.
func parseRow(s string) game {
	gameString := strings.Split(s, ":")
	gameID := strings.Split(gameString[0], " ")
	gameIDn := gameID[len(gameID)-1]
	gameIDint, err := strconv.Atoi(gameIDn)

	if err != nil {
		fmt.Printf("Couldn't convert game ID to int: %v", err)
		os.Exit(1)
	}

	// initialize a slice of string -> int maps to be filled with reveals for
	// this game
	var reveals []map[string]int

	// use the second part of the previously split string, separate reveals on
	// ';' and loop over them
	revealString := strings.Split(gameString[1], ";")
	for _, x := range revealString {
		// reveals themselves are separated by commas
		splitReveals := strings.Split(x, ",")
		// create a map for color -> count relations that we can fill and append
		// to 'reveals' later
		revealMap := make(map[string]int)

		for _, rev := range splitReveals {
			// within a reveal, our int and string are separated by a space. we
			// use strings.TrimSpace to get rid of any leading or following
			// spaces.
			intString := strings.Split(strings.TrimSpace(rev), " ")
			revealInt, _ := strconv.Atoi(intString[0])
			revealColor := intString[1]
			// now that we have an integer and a string, we can put it into our
			// revealMap created before the loop
			revealMap[revealColor] = revealInt
		}
		// all loops for this reveal set are done, append our map[string]int to
		// the parent 'reveals'
		reveals = append(reveals, revealMap)
	}
	// return a 'game', identified by gameIDint with all our 'reveals'
	var gameReturn = game{
		id:      gameIDint,
		reveals: reveals,
	}
	return gameReturn
}

// validateGame takes a 'game' and returns whether that game would have been
// possible or not based on hard-coded criteria (below)
func validateGame(g game) bool {
	// "The Elf would first like to know which games would have been possible if
	// the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?"
	max_red := 12
	max_green := 13
	max_blue := 14
	// check for impossibility, otherwise return true
	for _, reveal := range g.reveals {
		if reveal["red"] > max_red || reveal["green"] > max_green || reveal["blue"] > max_blue {
			return false
		}
	}
	return true
}

// getMaxCubeCount returns the largest amount of cubes present within a game,
// which is also the fewest amount of cubes required to make the game possible
func getMaxCubeCount(g game) map[string]int {
	max_red := 0
	max_green := 0
	max_blue := 0
	returnMap := make(map[string]int)

	for _, reveal := range g.reveals {
		// check for new max values
		if reveal["red"] > max_red {
			max_red = reveal["red"]
		}
		if reveal["green"] > max_green {
			max_green = reveal["green"]
		}
		if reveal["blue"] > max_blue {
			max_blue = reveal["blue"]
		}
	}

	returnMap["red"] = max_red
	returnMap["green"] = max_green
	returnMap["blue"] = max_blue

	return returnMap

}

// powerOfCubeCount returns the power of ints in the getMaxCubeCount return
func powerOfCubeCount(input map[string]int) int {
	sum := 1
	for _, val := range input {
		sum = val * sum
	}
	return sum
}
