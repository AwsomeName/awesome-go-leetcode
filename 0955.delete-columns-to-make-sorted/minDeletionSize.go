package problem0955

import "fmt"

func minDeletionSize(A []string) int {
	fmt.Println("new--------------------------", A)
	ans := 0
	cuts := make([]bool, len(A[0]))

	for i := range A[0] {
		tmp := true
		for j := 0; j < len(A)-1; j++ {
			if !cuts[j] && A[j][i] > A[j+1][i] {
				ans++
				fmt.Println("need move:", i, cuts[j], A[j][i], A[j+1][i])
				tmp = false
				break
			}
		}
		if tmp {
			for k := 0; k < len(A)-1; k++ {
				if A[k][i] < A[k+1][i] {
					cuts[k] = true
				}
			}
		}
	}
	return ans
}
