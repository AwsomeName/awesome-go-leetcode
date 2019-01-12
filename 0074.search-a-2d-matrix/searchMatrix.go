package problem0074

func searchMatrix(matrix [][]int, target int) bool {
	m := matrix
	if len(m) <= 0 {
		return false
	}
	rows, cols := len(m), len(m[0])
	if cols <= 0 {
		return false
	}
	idxI, idxJ := 0, 0
	for idxI < rows {
		// fmt.Println(idxI,idxJ,m[idxI],m[idxI][0])
		if m[idxI][0] > target {
			return false
		}
		if m[idxI][cols-1] < target {
			idxI++
			continue
		}
		for idxJ = 0; idxJ <= cols; idxJ++ {
			if m[idxI][idxJ] == target {
				return true
			} else if m[idxI][idxJ] > target {
				return false
			}
		}
		return false
	}
	return false
}
