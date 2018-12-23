package problem0962

func maxWidthRamp(A []int) int {
    var find func(A []int) int
    find = func(A []int) int{
        if len(A) < 2 {
            return 0
        }
        i,j := 0,len(A)-1
        for i<j {
            if A[i] <= A[j] {
                return j-i
            }else {
                return max( find(A[1:]), find(A[:len(A)-1]) )
            }
        }
        return 0
    }

    return find(A)
}

func max(a,b int) int{
    if a > b {
        return a
    }else {
        return b
    }
}
