package problem0717

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if n <= 1 {
		return true
	}
	if bits[n-2] == 0 {
		return true
	}
	cnt := 0
	for i := n - 2; i >= 0 && bits[i] == 1; i-- {
		cnt++
	}
	if cnt%2 == 0 {
		return true
	} else {
		return false
	}
}
