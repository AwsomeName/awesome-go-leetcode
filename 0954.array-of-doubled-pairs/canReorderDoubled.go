package problem0954

import "fmt"
import "sort"

func canReorderDoubled(A []int) bool {
	fmt.Println("-------------------")
	if len(A) == 0 {
		return true
	}
	sort.Ints(A)
	index := make([]bool, len(A))

	AMap := func(find int) (int, bool) {
		if find < 0 {
			find = find / 2
		} else {
			find = find * 2
		}
		left, right := 0, len(A)
		for left < right {
			mid := (left + right) / 2
			if !index[mid] && A[mid] == find {
				return mid, true
			} else if find > A[mid] {
				left = mid + 1
			} else if find < A[mid] {
				right = mid
			} else if find == A[mid] {
				left++
			}
		}
		return 0, false
	}

	for i := 0; i < len(A); i++ {
		if index[i] {
			continue
		}
		if tmp, ok := AMap(A[i]); ok {
			index[tmp] = true
			index[i] = true
			fmt.Println(tmp, i)
		} else {
			return false
		}
	}

	return true
}
