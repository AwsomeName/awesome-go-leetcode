package problem0054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0])==0 {
		return []int{}
	}

	total := len(matrix) * len(matrix[0])
	m := len(matrix)
	n := len(matrix[0])
	marks := make([][]bool, m)
	for i := range marks {
		marks[i] = make([]bool, n)
	}

	ans := make([]int, total)
	ans[0] = matrix[0][0]
	marks[0][0] = true
	direct, mm, nn := 0, 0, 0
	var next func() int
	next = func() int {
		switch direct {
		case 0:
			if nn+1 >= n || marks[mm][nn+1] {
				direct++
				return next()
			} else {
				nn++
				marks[mm][nn] = true
				return matrix[mm][nn]
			}
		case 1:
			if mm+1 >= m || marks[mm+1][nn] {
				direct++
				return next()
			} else {
				mm++
				marks[mm][nn] = true
				return matrix[mm][nn]
			}
		case 2:
			if nn-1 < 0 || marks[mm][nn-1] {
				direct++
				return next()
			} else {
				nn--
				marks[mm][nn] = true
				return matrix[mm][nn]
			}

		case 3:
			if mm-1 < 0 || marks[mm-1][nn] {
				direct = 0
				return next()
			} else {
				mm--
				marks[mm][nn] = true
				return matrix[mm][nn]
			}
		default:
			return -1
		}
	}

	for i := 1; i < total; i++ {
		ans[i] = next()
	}
	return ans
}
