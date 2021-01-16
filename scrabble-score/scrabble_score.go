package scrabble

import "strings"

var scrabblePoints = map[rune]int{}

func initScrabblePoints() {
	var letterValues = map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}

	for letters, value := range letterValues {
		for _, letter := range letters {
			scrabblePoints[letter] = value
		}
	}
}

func Score(s string) int {
	initScrabblePoints()

	points := 0
	for _, char := range strings.ToLower(s) {
		points += scrabblePoints[char]
	}

	return points
}
