package problem0062

func uniquePaths(m int, n int) int {
	if m <= 1 || n <= 1 {
		return min(m, n)
	}

	a := max(m, n)
	b := min(m, n)

	a--
	b--

	ans := CC(a+b, b) / CC(b, b)

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CC(a, b int) int {
	ans := a
	a--
	for i := 0; a > 0 && i < b-1; i++ {
		ans *= a
		a--
	}
	return ans
}
