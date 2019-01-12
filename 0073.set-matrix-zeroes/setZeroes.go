package problem0073

func setZeroes(matrix [][]int) {
	if len(matrix) <= 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)

	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := range matrix {
		for j := range matrix[0] {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}
