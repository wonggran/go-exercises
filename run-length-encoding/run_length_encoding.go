package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	var current byte
	currentCount := 0
	var encoded strings.Builder
	for i := 0; i < len(input); i++ {
		if i == 0 { // at the beginning there won't be a different rune but there will be one of them
			current = input[i]
			currentCount++
		} else { // at intermediate runes check if it's the same as current
			// if it is then simply count++ otherwise push the count and current onto string builder and
			// set new current and set its count to 1
			if input[i] != current {
				if currentCount > 1 {
					encoded.WriteString(fmt.Sprintf("%d", currentCount))
				}
				encoded.WriteByte(byte(current))
				current = input[i]
				currentCount = 1
			} else {
				currentCount++
			}
		}
	}

	// after the last rune has been processed it hasn't been written but current and its count are
	// correct

	if currentCount > 1 {
		encoded.WriteString(fmt.Sprintf("%d", currentCount))

	}
	encoded.WriteByte(byte(current))

	return encoded.String()
}

func RunLengthDecode(input string) string {
	if len(input) == 0 || len(input) == 1 {
		return input
	}

	byteInput := []byte(input)
	var decoded strings.Builder
	var currentCount strings.Builder

	for i := 0; i < len(byteInput); i++ {
		current := byteInput[i]
		if i == 0 {
			if unicode.IsNumber(rune(current)) {
				currentDigit, _ := strconv.Atoi(string(current))
				currentCount.WriteString(fmt.Sprintf("%d", currentDigit))
			} else {
				decoded.WriteByte(current)
			}
		} else {
			if unicode.IsNumber(rune(current)) {
				currentDigit, _ := strconv.Atoi(string(current))
				currentCount.WriteString(fmt.Sprintf("%d", currentDigit))
			} else {
				n, _ := strconv.Atoi(currentCount.String())
				if n > 1 {
					for j := 0; j < n; j++ {
						decoded.WriteByte(current)
					}
				} else {
					decoded.WriteByte(current)
				}
				currentCount.Reset()
			}

		}
	}

	return decoded.String()
}
