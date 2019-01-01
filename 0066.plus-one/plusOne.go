package problem0066

func plusOne(digits []int) []int {
	dLen := len(digits)
	ans := make([]int, dLen+1)
	tmp := 1

	for i := dLen; i >= 1; i-- {
		ans[i] = digits[i-1] + tmp
		if ans[i] >= 10 {
			tmp = 1
			ans[i] -= 10
		} else {
			tmp = 0
		}
	}

	if tmp == 1 {
		ans[0] = 1
		return ans
	}

	return ans[1:]
}
