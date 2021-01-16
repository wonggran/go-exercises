package acronym

import (
	"strings"
	"unicode"
)

/*
Appends the first rune of a - or _ surrounded field to abbr.
*/
func runeToAppendBy(splitted []string, abbr *strings.Builder) {
	for _, word := range splitted {
		if len(word) > 0 {
			firstRune := rune(word[0])
			if unicode.IsLetter(firstRune) {
				(*abbr).WriteRune(unicode.ToUpper(firstRune))
			}
		}

	}

}

func Abbreviate(s string) string {
	var abbr strings.Builder
	for _, field := range strings.Fields(s) {
		// Fields may contain - or _
		splitByDash := strings.Split(field, "-")
		splitByUnderscore := strings.Split(field, "_")

		fieldHasDash := len(splitByDash) > 1
		fieldHasUnderscore := len(splitByUnderscore) > 1

		if fieldHasDash {
			runeToAppendBy(splitByDash, &abbr)
		} else if fieldHasUnderscore {
			runeToAppendBy(splitByUnderscore, &abbr)
		} else {
			abbr.WriteRune(unicode.ToUpper(rune(field[0])))
		}
	}
	return abbr.String()
}
