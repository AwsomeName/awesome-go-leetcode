package problem1022

func smallestRepunitDivByK(K int) int {
    // fmt.Println(math.MaxInt64)

    x := 1

    if x % K == 0 {
        return 1
    }
    if K % 2 == 0 || K % 5 == 0 {
        return -1
    }

    cnt := 1
    for cnt <= K {
        x = (x*10 + 1) % K
        cnt ++
        if x == 0 {
            // fmt.Println(cnt, x, K)
            return cnt
        }
    }

    return -1
}
