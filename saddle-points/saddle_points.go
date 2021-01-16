package matrix

import (
	"strconv"
	"strings"
)

func (m Matrix) isSaddle(i, j int) bool {
	saddleEle := m[i][j]

	for k := 0; k < len(m); k++ {
		colEle := m[k][j]
		if saddleEle > colEle {
			return false
		}
	}

	for k := 0; k < len(m[i]); k++ {
		rowEle := m[i][k]
		if saddleEle < rowEle {
			return false
		}
	}

	return true
}

// Pair represents the row i and column j element of a 2D matrix.
type Pair struct {
	i int
	j int
}

type Matrix [][]int

// Saddle returns all the saddle points in this matrix.
func (m Matrix) Saddle() []Pair {
	saddlePairs := []Pair{}

	for i, _ := range m {
		for j, _ := range m[i] {
			if m.isSaddle(i, j) {
				saddlePairs = append(saddlePairs, Pair{i, j})
			}
		}
	}

	return saddlePairs
}

func parse(input string) Matrix {
	levelOne := strings.Split(input, "\n")

	levelTwo := [][]string{}
	for _, str := range levelOne {
		levelTwo = append(levelTwo, strings.Split(str, " "))
	}

	levelThree := [][]int{}
	for i, str := range levelTwo {
		levelThree = append(levelThree, make([]int, len(str)))
		for j, runeEle := range str {
			levelThree[i][j], _ = strconv.Atoi(string(runeEle))
		}
	}

	return levelThree
}

// New returns a pointer to a 2D integer matrix which is constructed from the input string.
func New(input string) (*Matrix, error) {
	matrix := parse(input)
	return &matrix, nil
}
