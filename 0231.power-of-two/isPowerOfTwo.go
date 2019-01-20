package problem0231

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}

	idx := 1
	for idx <= n {
		if idx == n {
			return true
		}
		idx *= 2
	}
	return false
}
