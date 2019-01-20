package problem0980

func uniquePathsIII(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	idxI1, idxJ1, idxI2, idxJ2 := 0, 0, 0, 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				idxI1, idxJ1 = i, j
			}
			if grid[i][j] == 2 {
				idxI2, idxJ2 = i, j
			}
		}
	}

	return move(grid, idxI1, idxJ1, idxI2, idxJ2)
}

func move(grid [][]int, idxI1, idxJ1, idxI2, idxJ2 int) int {
	if check(grid) {
		if (idxI1-idxI2)*(idxI1-idxI2)+(idxJ1-idxJ2)*(idxJ1-idxJ2) == 1 {
			return 1
		} else {
			return 0
		}
	}
	cnt := 0
	if idxI1 > 0 && grid[idxI1-1][idxJ1] == 0 {
		grid[idxI1-1][idxJ1] = 1
		cnt += move(grid, idxI1-1, idxJ1, idxI2, idxJ2)
		grid[idxI1-1][idxJ1] = 0
	}
	if idxI1 < len(grid)-1 && grid[idxI1+1][idxJ1] == 0 {
		grid[idxI1+1][idxJ1] = 1
		cnt += move(grid, idxI1+1, idxJ1, idxI2, idxJ2)
		grid[idxI1+1][idxJ1] = 0
	}
	if idxJ1 > 0 && grid[idxI1][idxJ1-1] == 0 {
		grid[idxI1][idxJ1-1] = 1
		cnt += move(grid, idxI1, idxJ1-1, idxI2, idxJ2)
		grid[idxI1][idxJ1-1] = 0
	}
	if idxJ1 < len(grid[0])-1 && grid[idxI1][idxJ1+1] == 0 {
		grid[idxI1][idxJ1+1] = 1
		cnt += move(grid, idxI1, idxJ1+1, idxI2, idxJ2)
		grid[idxI1][idxJ1+1] = 0
	}

	return cnt
}

func check(grid [][]int) bool {
	fmt.Println(grid)
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
