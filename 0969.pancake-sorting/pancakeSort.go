package problem0969

func pancakeSort(A []int) []int {
	res := make([]int, 0)

	find := func(i int) int {
		for j := 0; j < i; j++ {
			if A[j] == i {
				return j
			}
		}
		return -1
	}

	flip := func(index int) {
		for i := 0; i <= index/2; i++ {
			// fmt.Println(i,index-i,index)
			A[i], A[index-i] = A[index-i], A[i]
		}
	}

	for i := len(A); i >= 1; i-- {
		idx := find(i)
		if idx+1 != i {
			flip(idx)
			flip(i - 1)
			res = append(res, idx+1)
			res = append(res, i)
		}
	}

	return res
}
