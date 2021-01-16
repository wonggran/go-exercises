package transpose

func Transpose(input []string) []string {
	n := 0 // to be the number of rows needed
	for _, row := range input {
		if len(row) > n {
			n = len(row)
		}
	}
	processed := [][]byte{} // pad
	for i, row := range input {
		processed = append(processed, []byte{})
		processed[i] = []byte(row)
	}
	isTriangle := true
	for i := 0; i < len(input)-1; i++ {
		if len(input[i+1]) != len(input[i])+1 {
			isTriangle = false
			break
		}
	}
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if len(input[i]) < len(input[j]) {
				lenDiff := len(input[j]) - len(input[i])
				if isTriangle {
					lenDiff = n - len(input[i])
				}
				for k := 0; k < lenDiff; k++ {
					processed[i] = append(processed[i], ' ')
				}
			}
			break
		}

	}
	builder := [][]byte{} // each row
	for i := 0; i < n; i++ {
		builder = append(builder, []byte{})
	}
	for i := 0; i < len(processed); i++ {
		for j := 0; j < len(processed[i]); j++ {
			builder[j] = append(builder[j], processed[i][j])
		}
	}
	ret := []string{}
	for _, runeSlice := range builder {
		ret = append(ret, string(runeSlice))
	}
	return ret
}
