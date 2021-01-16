package spiralmatrix

// SpiralMatrix returns an n*n matrix containing a clockwise spiral of ascending integers.
func SpiralMatrix(n int) [][]int {
	spiral := make([][]int, n)
	for i := 0; i < n; i++ {
		spiral[i] = make([]int, n)
	}

	offset := 0
	topRow, bottomRow := 0, n-1
	leftCol, rightCol := 0, n-1
	for i := 1; i <= n*n; {
		for t := offset; t < n-offset-1; t++ {
			spiral[topRow][t] = i
			i++
		}
		topRow++

		for r := offset; r < n-offset-1; r++ {
			spiral[r][rightCol] = i
			i++
		}
		rightCol--

		for b := n - 1 - offset; b > offset; b-- {
			spiral[bottomRow][b] = i
			i++
		}
		bottomRow--

		for l := n - 1 - offset; l > offset; l-- {
			spiral[l][leftCol] = i
			i++
		}
		leftCol++

		offset++

		if i == n*n {
			mid := int(n / 2)
			spiral[mid][mid] = i
			i++
		}
	}

	return spiral
}
