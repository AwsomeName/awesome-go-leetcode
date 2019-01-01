package problem0059

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	marks := make([][]bool, n)
	for i := range marks {
		marks[i] = make([]bool, n)
	}

	idx := 2
	ans[0][0] = 1
	marks[0][0] = true

	direct, mm, nn := 0, 0, 0
	var next func()
	next = func() {
		switch direct {
		case 0:
			if nn+1 >= n || marks[mm][nn+1] {
				direct++
				next()
			} else {
				nn++
				marks[mm][nn] = true
				// return matrix[mm][nn]
				ans[mm][nn] = idx
				idx++
				return
			}
		case 1:
			if mm+1 >= n || marks[mm+1][nn] {
				direct++
				next()
			} else {
				mm++
				marks[mm][nn] = true
				// return matrix[mm][nn]
				ans[mm][nn] = idx
				idx++
				return
			}
		case 2:
			if nn-1 < 0 || marks[mm][nn-1] {
				direct++
				next()
			} else {
				nn--
				marks[mm][nn] = true
				// return matrix[mm][nn]
				ans[mm][nn] = idx
				idx++
				return
			}

		case 3:
			if mm-1 < 0 || marks[mm-1][nn] {
				direct = 0
				next()
			} else {
				mm--
				marks[mm][nn] = true
				// return matrix[mm][nn]
				ans[mm][nn] = idx
				idx++
				return
			}
		default:
			return
		}
	}

	for i := 1; i < n*n; i++ {
		next()
	}
	return ans
}
