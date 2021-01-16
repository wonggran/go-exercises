package pascal

func Triangle(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	rows := [][]int{{1}}

	for i := 1; i < n; i++ {
		prevRow := rows[len(rows)-1]
		prevRow = append([]int{0}, prevRow...) // pad for left, right
		prevRow = append(prevRow, 0)
		currentRow := []int{}
		for j := 0; j < len(prevRow)-1; j++ {
			currentRow = append(currentRow, prevRow[j]+prevRow[j+1])
		}
		rows = append(rows, currentRow)
	}

	return rows
}
