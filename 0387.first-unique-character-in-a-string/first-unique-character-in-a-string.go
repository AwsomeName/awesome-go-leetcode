package problem0387

func firstUniqChar(s string) int {
	marks := make(map[rune]int)

	for _, x := range s {
		marks[x]++
	}

	for i, x := range s {
		if marks[x] == 1 {
			return i
		}
	}
	return -1
}
