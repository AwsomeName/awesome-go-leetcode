package problem0955

import "fmt"

func minDeletionSize(A []string) int {
	fmt.Println("-------------------")
	index := make([]bool, len(A[0]))
	old := 0
    res := 0
	for {
        old = res
		for i := 0; i < len(A)-1; i++ {
			lOrder(A[i], A[i+1], index, 0,&res)
            fmt.Println("A:",A[i],A[i+1],"index:",index,res)
		}
		if old == res || res >= len(A[0]) {return res}
	}
}

func lOrder(A string, B string, index []bool, start int,res *int) {
	if start > len(A)-1 {
		return
	}
	if index[start] {
		lOrder(A, B, index, start+1,res)
	}
	if A[start] > B[start] {
        if !index[start] {
	    	index[start] = true
            *res ++
        }
		lOrder(A, B, index, start+1,res)
	} else if A[start] == B[start] {
		lOrder(A, B, index, start+1,res)
	}
}
