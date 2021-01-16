package grains

import "errors"

func Square(square int) (uint64, error) {
	if !(square >= 1 && square <= 64) {
		return 0, errors.New("Square numbers are positive integers between 1 and 64.")
	}

	return 1 << (square - 1), nil
}

/* Returns the total number of grains for a 64 tile chess board with each successive
tile having double the number of grains of the previous tile. */
func Total() uint64 {
	n := 64
	var total uint64

	for i := 1; i <= n; i++ {
		grains, _ := Square(i)
		total += grains
	}

	return total
}
