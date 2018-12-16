package problem0957

import "fmt"

func prisonAfterNDays(cells []int, N int) []int {
	N = N%14 + 14
	tmp := make([]int, len(cells))
	tmp = dcopy(cells)
	tmp[0], tmp[7] = 0, 0
	for i := 0; i < N; i++ {
		for j := 1; j <= 6; j++ {
			if cells[j-1] == cells[j+1] {
				tmp[j] = 1
			} else {
				tmp[j] = 0
			}
		}
		cells = dcopy(tmp)
		fmt.Println("cells:", cells)
	}
	return cells
}
func dcopy(cells []int) []int {
	tmp := make([]int, len(cells))
	for i, x := range cells {
		tmp[i] = x
	}
	return tmp
}
