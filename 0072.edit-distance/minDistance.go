package problem0072

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	if len1 == 0 {
		return len2
	}

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := range word2 {
		dp[0][i+1] = i + 1
	}

	for i := range word1 {
		dp[i+1][0] = i + 1
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len1][len2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
