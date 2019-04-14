package problem1012

func bitwiseComplement(N int) int {
    if N == 0 { return 1 }

    cnt1 := make([]int,0)
    // cnt2 := make([]int,0)

    for N > 0 {
        cnt1 = append(cnt1, 1- ( N % 2 ) )
        N = N / 2
    }

    ans := 0
    base := 1
    for _,x := range cnt1 {
        ans += (x * base)
        base = base * 2
    }

    return ans

}
