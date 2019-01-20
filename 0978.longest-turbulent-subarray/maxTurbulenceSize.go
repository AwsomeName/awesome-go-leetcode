package problem0978

func maxTurbulenceSize(A []int) int {
	if len(A) == 0 {
		return 0
	}

	cnt := 0
	flags := make([]int, len(A))

	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			flags[i] = 1
		} else if A[i] > A[i+1] {
			flags[i] = 2
		}
	}

	fmt.Println(flags)
	tmp := 0
	for i := 0; i < len(flags)-1; i++ {
		if flags[i] == flags[i+1] || flags[i] == 0 {
			if flags[i] != 0 {
				tmp++
			}
			cnt = max(cnt, tmp)
			tmp = 0
			continue
		}
		fmt.Println(i, tmp)
		tmp++
	}
	cnt = max(cnt, tmp)
	return cnt + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
