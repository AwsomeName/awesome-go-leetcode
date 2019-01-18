package problem0089

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	ans := make([]int, 1<<uint(n))
	for i := 0; i < 1<<uint(n); i++ {
		ans[i] = i ^ (i >> 1)
	}

	return ans
}
