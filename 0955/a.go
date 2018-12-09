package problem0955

import "fmt"

func minDeletionSize(A []string) int {
	fmt.Println("-------------------")
	res := 0
	for i := 0; i < len(A[0]); i++ {
		cnt := 0
		for j := 0; j < len(A)-1; j++ {
			tmp := A[j][i:]
			tmp2 := A[j+1][i:]
			fmt.Println("tmp:", tmp, "tmp2:", tmp2, "res:", lOrder(tmp, tmp2))
			if !lOrder(tmp, tmp2) {
				res++
				break
			}
			cnt++
		}
		if cnt == len(A)-1 {
			return res
		}
	}
	return res
}

func lOrder(A string, B string) bool {
	for i := 0; i < len(A); i++ {
		if A[i] < B[i] {
			return true
		} else if A[i] > B[i] {
			return false
		}
	}
	return true
}
