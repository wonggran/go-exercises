package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Input strings lengths differ!")
	}

	dist := 0
	for indA, byteA := range a {
		byteB := b[indA]
		if byteB != byte(byteA) {
			dist++
		}
	}

	return dist, nil
}
