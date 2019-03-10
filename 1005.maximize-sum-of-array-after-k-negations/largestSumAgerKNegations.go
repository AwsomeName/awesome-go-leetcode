package problem1005

func largestSumAfterKNegations(A []int, K int) int {
    sort.Ints(A)
    fmt.Println(A)
    if A[0] == 0 {
        return sumA(A)
    }

    i := 0
    for i < len(A) && A[i] < 0 && K > 0 {
        A[i] = -A[i]
        K --
        i++
    }
    // fmt.Println(A)
    if K == 0 {
        return sumA(A)
    }
    sort.Ints(A)
    // fmt.Println(A)
    for K>0 {
        K--
        A[0] = -A[0]
    }
    return sumA(A)

}


func sumA(A []int) int{
    ans := 0
    for _,x := range A{
        ans += x
    }
    return ans
}
