package problem0977

func sortedSquares(A []int) []int {
	if len(A) == 0 {
		return []int{}
	}

	ans := make([]int, len(A))

	for i, x := range A {
		ans[i] = x * x
	}
	sort.Ints(ans)
	return ans
}
