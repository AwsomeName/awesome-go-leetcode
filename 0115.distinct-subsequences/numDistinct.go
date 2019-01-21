package problem0115

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	ans := make([][]int, n+1)
	for i := range ans {
		ans[i] = make([]int, m+1)
	}
	for j := 0; j <= m; j++ {
		ans[0][j] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if t[i-1] == s[j-1] {
				ans[i][j] = ans[i-1][j-1] + ans[i][j-1]
			} else {
				ans[i][j] = ans[i][j-1]
			}
		}
		// fmt.Println(ans[i])
	}
	return ans[n][m]
}
