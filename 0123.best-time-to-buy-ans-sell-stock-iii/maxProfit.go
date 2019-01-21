package problem0123

func maxProfit(prices []int) int {
	cur := 0
	sum := 0
	cur2 := 0
	sum2 := 0
	for i := 1; i < len(prices); i++ {
		cur = max(0, cur+prices[i]-prices[i-1])
		sum = max(cur, sum)
		cur2 = max(sum, cur2+prices[i]-prices[i-1])
		sum2 = max(cur2, sum2)
	}

	return sum2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
