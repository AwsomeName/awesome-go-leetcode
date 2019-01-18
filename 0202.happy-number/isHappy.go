package problem0202

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	if n == 2 {
		return false
	}

	sum := sq(n)

	record := make(map[int]bool)

	for {
		if sum == 1 {
			return true
		}
		if record[sum] {
			return false
		} else {
			record[sum] = true
		}
		sum = sq(sum)
	}

	return true
}

func sq(a int) int {
	sum := 0
	for a > 0 {
		sum += (a % 10) * (a % 10)
		a = a / 10
	}
	return sum
}
