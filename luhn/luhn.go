package luhn

import (
	"errors"
	"math"
	"strconv"
	"unicode"
)

func CountDigits(num string) (int, error) {
	n := 0

	for _, char := range num {
		if !unicode.IsDigit(char) && !unicode.IsSpace(char) {
			return -1, errors.New("Invalid byte in sequence.")
		}

		if !unicode.IsSpace(char) {
			n++
		}
	}

	return n, nil
}

func CreateDigits(num string, n int) []int {
	digits := make([]int, n)
	idx := 0

	for _, char := range num {
		dig, err := strconv.Atoi(string(char))
		if err == nil {
			digits[idx] = dig
			idx++
		}
	}

	return digits
}

func IsValidLuhn(digs []int, n int) bool {
	sum := 0

	/* Zero indexing and counting starting from the right has every other element
	in an odd number of digits with an odd index and even number of digits with
	an even index. */
	everyOther := 1

	if int(math.Mod(float64(n), 2)) == 0 {
		everyOther = 0
	}

	for idx := n - 1; idx >= 0; idx-- {
		dig := digs[idx]
		// Double every other digit.

		if int(math.Mod(float64(idx), 2)) == everyOther {
			doubled := dig * 2

			if doubled > 9 {
				doubled -= 9
			}

			sum += doubled
		} else {
			sum += dig
		}
	}

	return int(math.Mod(float64(sum), 10)) == 0
}

func Valid(num string) bool {
	n, err := CountDigits(num)

	if n < 2 || err != nil {
		return false
	}

	digs := CreateDigits(num, n)

	return IsValidLuhn(digs, n)
}
