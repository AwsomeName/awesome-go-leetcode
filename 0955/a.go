package problem0955

import "fmt"

func minDeletionSize(A []string) int {
	fmt.Println("-------------------")
	index := make([]bool, len(A[0]))
	for i := 0; i < len(A)-1; i++ {
		lOrder(A[i], A[i+1], index, 0)
	}
	res := 0
	for _, x := range index {
		if x {
			res++
		}
	}
	return res
}

func lOrder(A string, B string, index []bool, start int) {
	if start > len(A)-1 {
		return
	}
	if index[start] {
		lOrder(A, B, index, start+1)
	}
	if A[start] > B[start] {
		index[start] = true
		lOrder(A, B, index, start+1)
	} else if A[start] == B[start] {
		lOrder(A, B, index, start+1)
	}
}
