package problem0091

func numDecodings(s string) int {
	dp := make([]int, len(s))
	if len(s) <= 0 {
		return len(s)
	}
	// if len(s) == 1 && s[0] == '0' { return 0 }
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dp[0] = 1
	if s[1] != '0' && cal(s[:2]) <= 26 {
		dp[1] = 2
	} else if s[1] == '0' && s[0] > '2' {
		// fmt.Println(dp)
		return 0
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(s); i++ {
		if s[i] != '0' {
			dp[i] += dp[i-1]
			if x := cal(s[i-1 : i+1]); s[i-1] != '0' && x <= 26 {
				dp[i] += dp[i-2]
			}
		} else {
			if s[i-1] == '0' || s[i-1] > '2' {
				fmt.Println(dp)
				return 0
			}
			dp[i] = dp[i-2]
		}
	}

	return dp[len(s)-1]
}

func cal(s string) int {
	if len(s) != 2 {
		return 0
	}
	return int((s[0]-'0')*10 + s[1] - '0')
}
