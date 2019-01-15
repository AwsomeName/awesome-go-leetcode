package problem0077

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	cmb(&ans, []int{}, n, k, 1, 0)
	return ans
}

func cmb(ans *[][]int, tmp []int, n, k int, sk, tk int) {
	if tk == k {
		*ans = append(*ans, tmp)
		return
	}

	for i := sk; i <= n; i++ {
		t := make([]int, len(tmp))
		copy(t, tmp)
		t = append(t, i)
		cmb(ans, t, n, k, i+1, tk+1)
	}
}
