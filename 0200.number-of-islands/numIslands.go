package problem0200

func numIslands(grid [][]byte) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	cnt := 0
	mask := make([][]bool, len(grid))
	for i := range mask {
		mask[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !mask[i][j] {
				cnt++
				FindAndMask(grid, mask, i, j)
			}
		}
	}

	return cnt
}

func FindAndMask(gird [][]byte, mask [][]bool, i, j int) {
	// if mask[i][j] {return}
	mask[i][j] = true
	if i+1 < len(gird) && gird[i+1][j] == '1' && !mask[i+1][j] {
		FindAndMask(gird, mask, i+1, j)
	}
	if i-1 >= 0 && gird[i-1][j] == '1' && !mask[i-1][j] {
		FindAndMask(gird, mask, i-1, j)
	}
	if j+1 < len(gird[0]) && gird[i][j+1] == '1' && !mask[i][j+1] {
		FindAndMask(gird, mask, i, j+1)
	}
	if j-1 >= 0 && gird[i][j-1] == '1' && !mask[i][j-1] {
		FindAndMask(gird, mask, i, j-1)
	}

	// mask[i][j] = true
}
