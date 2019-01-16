package problem0121

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	//minP, maxP := prices[0],prices[1]
	sum := 0
	t := 0
	for i := 1; i < len(prices); i++ {
		sum = max(0, sum+prices[i]-prices[i-1])
		t = max(t, sum)
	}
	prices[0] = 0
	return t
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// public int maxProfit(int[] prices) {
//       int maxCur = 0, maxSoFar = 0;
//       for(int i = 1; i < prices.length; i++) {
//           maxCur = Math.max(0, maxCur += prices[i] - prices[i-1]);
//           maxSoFar = Math.max(maxCur, maxSoFar);
//       }
//       return maxSoFar;
//   }
