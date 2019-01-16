package problem0168

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	idx := make(map[int]rune)
	for i, j := 'A', 1; j <= 26; j++ {
		idx[j] = i
		i++
	}
	idx[0] = 'Z'
	ans := ""
	for n >= 1 {
		ans = string(idx[n%26]) + ans
		if n%26 == 0 {
			n = n/26 - 1
		} else {
			n /= 26
		}
	}
	return ans
}
