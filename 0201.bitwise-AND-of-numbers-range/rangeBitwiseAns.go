package problem0201

func rangeBitwiseAnd(m int, n int) int {
	ans := m
	for i := m + 1; i <= n; i++ {
		ans = ans & i
		if ans == 0 {
			return 0
		}
	}
	return ans
}
