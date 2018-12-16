package problem0955

import "fmt"

func minDeletionSize(A []string) int {
	fmt.Println("new--------------------------", A)
    m,n := len(A),len(A[0])
    dp := make([]int, n)
    for i:= range dp{
        dp[i] = 1
    }

    for i:= 1; i < n ; i ++{
        for j := 0; j < i ; j ++{
            tmp := false
            for k:=0; k<m; k ++{
                if A[k][j] > A[k][i] {
                    tmp = false
                    break
                }
                tmp = true
            }
            if tmp {
                dp[i] = max(dp[i],dp[j]+1)
            }
        }
    }
    res := 0
    for _,x := range dp{
        if res < x {
            res = x
        }
    }
    return n-res
}


func max(a,b int) int{
    if a < b {
        return b
    }else {
        return a
    }
}
