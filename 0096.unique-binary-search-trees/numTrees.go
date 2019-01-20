package problem0096

// func numTrees(n int) int {
//     if n <= 1 {
//         return 1
//     }
//     cnt := 0
//     for i:= 1; i<=n; i++{
//         cnt+=(1 * numTrees(i-1) * numTrees(n-i)  )
//     }
//     return cnt
// }

func numTrees(n int) int {
	dp := make([]int, n+2)
	dp[0], dp[1] = 1, 1
	for k := 2; k <= n; k += 1 {
		for i := 0; i <= k-1; i += 1 {
			dp[k] += dp[i] * dp[k-1-i]
		}
	}
	return dp[n]
}
