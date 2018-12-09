package problem0953

import "fmt"

func isAlienSorted(words []string, order string) bool {
	fmt.Println("-------------------")
	orderMap := make(map[byte]int, len(order))
	for i := range order {
		orderMap[order[i]] = i
	}

	cmp := func(A string, B string) bool {
		i := 0
		for i < len(A) && i < len(B) {
			if A[i] == B[i] {
				i++
				continue
			} else if orderMap[A[i]] > orderMap[B[i]] {
				return false
			} else {
				return true
			}
		}
		if len(A) > len(B) {
			return false
		} else {
			return true
		}
	}

	for i := 0; i < len(words)-1; i++ {
		if !cmp(words[i], words[i+1]) {
			return false
		}
	}
	return true
}
