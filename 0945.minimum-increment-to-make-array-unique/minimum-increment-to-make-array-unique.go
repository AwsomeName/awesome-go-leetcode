package problem0945

import "sort"
import "fmt"

func minIncrementForUnique(A []int) int {
	if len(A) < 2 {
		return 0
	}
	sort.Ints(A)
	cnt := 0
	for i := 1; i < len(A); i++ {
		fmt.Println(A[i])
		for A[i] <= A[i-1] {
			cnt++
			A[i]++
		}
	}
	return cnt
}
