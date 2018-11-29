package problem0009

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	s := strconv.Itoa(x)
	for i := range s {
		if i >= len(s)/2 {
			break
		}
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
