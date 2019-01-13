package problem0118

func generate(numRows int) [][]int {
	idx := numRows
	if idx == 0 {
		return [][]int{}
	}
	if idx == 1 {
		return [][]int{{1}}
	}
	if idx == 2 {
		return [][]int{{1}, {1, 1}}
	}
	a := make([][]int, idx)
	for i := range a {
		a[i] = make([]int, i+1)
	}
	a[0][0] = 1
	a[1][0], a[1][1] = 1, 1

	for i := 2; i < idx; i++ {
		a[i][0] = 1
		for j := 1; j <= i-1; j++ {
			a[i][j] = a[i-1][j-1] + a[i-1][j]
		}
		a[i][i] = 1
	}
	return a
}
