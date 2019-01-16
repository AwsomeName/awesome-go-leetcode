package problem0171

func titleToNumber(s string) int {
	if len(s) == 0 {
		return 0
	}
	idx := make(map[rune]int)
	for i, j := 'A', 1; i <= 'Z'; i++ {
		idx[i] = j
		j++
	}

	ans := 0
	for _, x := range s {
		ans = ans*26 + idx[x]
	}

	return ans
}
