package problem0122

func maxProfit(prices []int) int {
	ps := prices
	if len(ps) < 2 {
		return 0
	}

	sum := 0
	for i := 1; i < len(ps); i++ {
		tmp := ps[i] - ps[i-1]
		if tmp > 0 {
			sum += tmp
		}
	}
	return sum
}
