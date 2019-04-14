package problem1020

func canThreePartsEqualSum(A []int) bool {
    if len(A) <= 2 {
        return false
    }
    i, j := 1,2
    sum0, sum1, sum2 := A[0], A[1], sum(A[j:])
    if sum0 == sum1 && sum1 == sum2 {
        return true
    }


    for i < len(A) -1 {
        sum0 = sum(A[0:i])
        sum1 = A[i]
        sum2 = sum(A[i+1:])
        if sum0 == sum1 && sum1 == sum2 {
            return true
        }
        for j :=i+1; j < len(A); j++{
            sum1 = sum1 + A[j]
            sum2 = sum2 - A[j]
            if sum0 == sum1 && sum1 == sum2 {
                return true
            }
        }
        i++
    }
    return false
}

func sum(A []int) int{
    if len(A)==0 {
        return 0
    }
    ans := 0
    for _,x := range A {
        ans += x
    }
    return ans
}
