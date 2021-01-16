package raindrops

import (
	"math"
	"strconv"
	"strings"
)

func Convert(n int) string {
	var ret strings.Builder
	input := float64(n)
	isMultipleFlag := false

	if math.Mod(input, 3) == 0 {
		ret.WriteString("Pling")
		isMultipleFlag = true
	}

	if math.Mod(input, 5) == 0 {
		ret.WriteString("Plang")
		isMultipleFlag = true

	}

	if math.Mod(input, 7) == 0 {
		ret.WriteString("Plong")
		isMultipleFlag = true

	}

	if !isMultipleFlag {
		return strconv.Itoa(n)
	}

	return ret.String()
}
