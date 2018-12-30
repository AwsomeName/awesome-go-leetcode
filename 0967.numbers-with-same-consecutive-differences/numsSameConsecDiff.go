package problem0967

func numsSameConsecDiff(N int, K int) []int {
	ans := make([]int, 0)

	// var genNum func(now, N, K int, ans *[]int)

	if N == 1 {
		x := 0
		for x < 10 {
			tmp := x
			ans = append(ans, tmp)
			x++
		}
		return ans
	}

	genNum(0, N, K, 0, &ans)

	return ans

}

func genNum(now, N, K, tmp int, ans *[]int) {
	if now == 0 {
		for i := 1; i < 10; i++ {
			// mid := []int{}
			// x := i
			// mid = append(mid, x)
			genNum(1, N, K, i, ans)
		}

	} else if now < N {
		x := tmp % 10
		if x+K < 10 {
			genNum(now+1, N, K, tmp*10+x+K, ans)
		}
		if x-K >= 0 && K != 0 {
			genNum(now+1, N, K, tmp*10+x-K, ans)
		}

	} else if now == N {
		*ans = append(*ans, tmp)
	}
}
