package reverse

func Reverse(s string) string {
	if len(s) > 0 {
		var sSlice []rune = []rune(s)
		front, back := 0, len(sSlice)-1

		for front < back {
			tmp := sSlice[front]
			sSlice[front] = sSlice[back]
			sSlice[back] = tmp
			front++
			back--
		}

		return string(sSlice)
	}

	return s
}
