package problem0962

func maxWidthRamp(A []int) int {
    mark := make(map[int]bool)
    idx := make([][]int,50001)
    for i:= range idx{
        idx[i] = make([]int,2)
    }
    for i,x := range A{
        if mark[x]{
            idx[x][1] = i
        }else{
            mark[x] = true
            idx[x][0] = i
            idx[x][1] = i
        }
    }
    ss := 50002
    ans := 0
    for i :=0; i<50001; i++{
        if mark[i]{
            ss = min(ss,idx[i][0])
            ans = max(ans,idx[i][1]-ss)
        }
    }
    return ans
}

func max (a,b int) int{
    if a < b {
        return b
    }else {
        return a
    }
}

func min (a,b int) int{
    if a < b {
        return a
    }else {
        return b
    }
}
