package problem1021

func maxScoreSightseeingPair(A []int) int {
    // maxI := make([]int, 0)
    if len(A) < 1 {
        return 0
    }
    tmpMax := A[0]
    max := 0
    for i:=1; i<len(A); i++{
        tmpSum := tmpMax - 1 + A[i]
        max = bigger(max, tmpSum)
        tmpMax = bigger(tmpMax-1, A[i])
    }
    return max
}

func bigger(a, b int) int{
    if a > b {
        return a
    }
    return b
}
