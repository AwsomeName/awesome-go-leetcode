package problem0120

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	t := triangle

	for i := rows - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			t[i][j] = t[i][j] + min(t[i+1][j], t[i+1][j+1])
		}
	}
	return t[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
