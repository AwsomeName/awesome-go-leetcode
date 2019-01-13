package problem0119

func getRow(rowIndex int) []int {
	idx := rowIndex + 1
	if idx == 0 {
		return []int{}
	}
	if idx == 1 {
		return []int{1}
	}
	if idx == 2 {
		return []int{1, 1}
	}

	a, b := make([]int, idx), make([]int, idx)
	a[0] = 1
	b[0], b[1] = 1, 1
	flag := true

	for i := 3; i <= idx; i++ {
		for j := 1; j < i-1; j++ {
			if flag {
				a[j] = b[j-1] + b[j]
				a[i-1] = 1
				// flag = false
			} else {
				b[j] = a[j-1] + a[j]
				b[i-1] = 1
				// flag = true
			}
		}
		flag = !flag
	}

	if flag {
		return b
	}
	return a
}
