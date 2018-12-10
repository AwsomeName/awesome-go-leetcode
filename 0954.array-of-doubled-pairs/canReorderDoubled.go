package problem0954

import "fmt"
import "sort"

func canReorderDoubled(A []int) bool {
	fmt.Println("-------------------")
	if len(A) == 0 {
		return true
	}
	sort.Ints(A)
	AMap := make(map[int]int)
	for i, x := range A {
		AMap[x] = i
	}
	index := make([]bool, len(A))

	for i := 0; i < len(A); i++ {
		if index[i] {
			continue
		}
		if A[i] > 0 {
			if tmp, ok := AMap[2*A[i]]; ok {
				index[tmp] = true
				index[i] = true
			} else {
				return false
			}
		} else {
			if tmp, ok := AMap[A[i]/2]; ok {
				index[tmp] = true
				index[i] = true
			} else {
				return false
			}
		}
	}

	return true
}
