package alphametics

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// creation of a rune matrix used to create a matrix of variables

func maxColumns(addends []string, answer string) int {
	max := 0

	for _, addend := range addends {
		if len(addend) > max {
			max = len(addend)
		}
	}

	if len(answer) > max {
		max = len(answer)
	}

	return max
}

func rows(addends []string) int {
	return len(addends) + 1
}

func createRuneMatrix(input string) [][]rune {
	adds := input[:strings.Index(input, "==")-1]
	answer := input[strings.Index(input, "==")+3:]

	addends := strings.Split(adds, " + ")

	numCols := maxColumns(addends, answer)
	numRows := rows(addends)

	matrix := [][]rune{}

	for i := 0; i < numRows; i++ {
		matrix = append(matrix, make([]rune, numCols))
	}

	for i, addend := range addends {
		for j, r := range addend {
			matrix[i][j] = r
		}
	}

	for j, r := range answer {
		matrix[numRows-1][j] = r
	}

	return matrix
}

// A variable represents a letter

type variable struct {
	currentValue int
	character    rune // the actual letter this variable represents
}

func (v variable) String() string {
	return fmt.Sprintf("%U", v.character)
}

// Creation of variable matrix

var vars []*variable

func createVariableMatrix(runeMatrix [][]rune) [][]*variable {
	matrix := [][]*variable{}

	runeToVariable := map[rune]*variable{}

	for i, _ := range runeMatrix { // [[a b c ] [d e] [f g h]]
		matrix = append(matrix, make([]*variable, len(runeMatrix[i])))
		for j, _ := range runeMatrix[i] {
			currentRune := runeMatrix[i][j]
			if unicode.IsLetter(currentRune) {
				if existingVariable, seen := runeToVariable[currentRune]; !seen {
					newVariable := &variable{character: currentRune}
					matrix[i][j] = newVariable       // the matrix stores references to variables
					vars = append(vars, newVariable) // the slice stores the actual variable
					// error source: the value of newVariable was being appended to vars not a reference
					runeToVariable[currentRune] = newVariable
				} else {
					// point to the existing variable
					matrix[i][j] = existingVariable
				}
			}
		}
	}
	return matrix
}

// Brute force a solution with the constraint that: all digits are unique and
// numbers don't lead with zero.

func variablesAreUnique() bool {
	digits := map[int]bool{}
	for _, v := range vars {
		if _, seen := digits[v.currentValue]; seen {
			return false
		}
		digits[v.currentValue] = true
	}
	return true
}

func createNumber(row []*variable) (int, error) {
	var strNumBuilder strings.Builder
	for i, val := range row {
		if val != nil {
			digit := fmt.Sprintf("%d", row[i].currentValue)
			strNumBuilder.WriteString(digit)
		}
	}

	strNum := strNumBuilder.String()
	if strNum[0] == '0' {
		return 0, errors.New("Leading digit is zero.")
	}

	num, err := strconv.Atoi(strNum)

	if err != nil {
		fmt.Println("err", err)
	}

	return num, nil
}

func formSolution() map[string]int {
	soln := map[string]int{}

	for _, v := range vars { // vars should be singly existing
		soln[string(v.character)] = v.currentValue
	}

	return soln
}

// A factorial solution!
func bruteForce(variableMatrix [][]*variable) (map[string]int, error) {
	n := len(vars)
	for i := 0; i < int(math.Pow(10, float64(n))); i++ { // All assignments possible
		assignment := fmt.Sprintf("%0*d", n, i)
		seenDigits := map[int]bool{}
		skipAssignment := false

		for j := 0; j < len(vars); j++ { // Make an assignment to all the variables
			digit, _ := strconv.Atoi(string(assignment[j]))
			if _, seen := seenDigits[digit]; !seen {
				seenDigits[digit] = true
				vars[j].currentValue = digit
			} else {
				break // Unique constraint
				skipAssignment = true
			}
		}

		if skipAssignment { // Skip this possible assignment because digits are not unique
			continue
		}

		// Digits are unique, now sum up the addends by creating the corresponding numerical value
		// for a row of []*variable
		addendLeadingZeroFlag := false
		sum := 0
		for i := 0; i < len(variableMatrix)-1; i++ {
			addend, errLeadingZero := createNumber(variableMatrix[i])

			if errLeadingZero != nil {
				addendLeadingZeroFlag = true
				break
			}

			sum += addend
		}

		numResult, errLeadingZero := createNumber(variableMatrix[len(variableMatrix)-1])

		if errLeadingZero != nil || addendLeadingZeroFlag {
			continue
		} // All numbers are now valid, finally check if equation is valid

		if sum == numResult && variablesAreUnique() {
			return formSolution(), nil
		}
	}

	return map[string]int{}, errors.New("No unique assignment solution with non-zero leading digits exists.")
}

func Solve(input string) (map[string]int, error) {
	runeMatrix := createRuneMatrix(input)
	variableMatrix := createVariableMatrix(runeMatrix)
	return bruteForce(variableMatrix)
}
