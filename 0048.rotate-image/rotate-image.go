package problem0048

func rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			m[i][j], m[n-j-1][i], m[n-i-1][n-j-1], m[j][n-i-1] = m[n-j-1][i], m[n-i-1][n-j-1], m[j][n-i-1], m[i][j]
		}
	}
}
