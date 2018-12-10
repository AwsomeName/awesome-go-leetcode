package problem0954

import "fmt"
import "sort"

func canReorderDoubled(A []int) bool {
	fmt.Println("--------", A, "-----------")
	if len(A) == 0 {
		return true
	}
	sort.Ints(A)
	AMap := make(map[int]int)
	for _, x := range A {
		AMap[x]++
	}

	for i := 0; i < len(A); i++ {
		if AMap[A[i]] <= 0 {
			continue
		}
		X := 0
		if A[i] < 0 {
			X = A[i] / 2
		} else {
			X = A[i] * 2
		}
		if tmp, ok := AMap[X]; ok {
			if tmp > 0 {
				AMap[A[i]]--
				AMap[X]--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
