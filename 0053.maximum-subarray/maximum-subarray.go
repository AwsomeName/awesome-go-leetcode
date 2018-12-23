package problem0053

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = max(dp[i-1]+nums[i], 0+nums[i])
		} else {
			dp[i] = nums[i]
		}
	}
	max := math.MinInt32
	for _, x := range dp {
		if max < x {
			max = x
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
