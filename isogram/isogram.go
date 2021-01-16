package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	lowerCaseInput := strings.ToLower(s)
	for i, char := range lowerCaseInput {
		if unicode.IsLetter(char) && strings.ContainsRune(lowerCaseInput[i+1:], char) {
			return false
		}
	}

	return true
}
