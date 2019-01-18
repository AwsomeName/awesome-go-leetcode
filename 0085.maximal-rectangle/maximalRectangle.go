package problem0085

func maximalRectangle(matrix [][]byte) int {
	mx := matrix
	m := len(mx)
	if m == 0 {
		return 0
	}
	n := len(mx[0])
	if m == 0 || n == 0 {
		return 0
	}
	left, right, height := make([]int, n), make([]int, n), make([]int, n)
	for i := range right {
		right[i] = n
	}
	ans := 0
	for i := 0; i < m; i++ {
		cur_left, cur_right := 0, n
		for j := 0; j < n; j++ {
			if mx[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		for j := 0; j < n; j++ {
			if mx[i][j] == '1' {
				left[j] = max(left[j], cur_left)
			} else {
				left[j] = 0
				cur_left = j + 1
			}

			if mx[i][n-j-1] == '1' {
				right[n-j-1] = min(right[n-j-1], cur_right)
			} else {
				right[n-j-1] = n
				cur_right = n - j - 1
			}
		}

		for j := 0; j < n; j++ {
			ans = max(ans, (right[j]-left[j])*height[j])
			// fmt.Println(left[j], right[j], height[j],ans,(right[j]-left[j])*height[j])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
