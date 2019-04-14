package problem1029

func prefixesDivBy5(A []int) []bool {
    if len(A) == 0 {
        return []bool{}
    }

    tmp := A[0]
    ans := make([]bool, len(A))
    if tmp == 1 {
        ans[0] = false
    }else{
        ans[0] = true
    }

    for i := 1; i < len(A); i++ {
        tmp = tmp * 2 + A[i]
        tmp = tmp % 10
        if tmp % 5 == 0 {
            ans[i] = true
        }else {
            ans[i] = false
        }
    }
    // fmt.Println( tmp)
    return ans
}
