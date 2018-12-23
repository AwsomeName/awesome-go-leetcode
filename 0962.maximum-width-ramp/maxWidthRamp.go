package problem0962

func maxWidthRamp(A []int) int {
	stk := make([]int, 0, 50001)
	for i, x := range A {
		if len(stk) == 0 || x < A[stk[len(stk)-1]] {
			stk = append(stk, i)
		}
	}
	fmt.Println(stk)
	res := 0
	for i := len(A) - 1; i >= 0; i-- {
		for len(stk) > 0 && A[stk[len(stk)-1]] <= A[i] {
			res = max(res, i-stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
