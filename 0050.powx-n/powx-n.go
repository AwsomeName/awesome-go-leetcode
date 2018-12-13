package problem0050

func myPow(x float64, n int) float64 {
	if x == 0 || x == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	var cnt func(x float64, n int) float64
	cnt = func(x float64, n int) float64 {
		if n == 2 {
			return x * x
		}
		if n == 1 {
			return x
		}
		if n%2 == 0 {
			return cnt(x, n/2) * cnt(x, n/2)
		} else {
			return cnt(x, n/2) * cnt(x, n/2) * x
		}
	}
	if n >= 0 {
		return cnt(x, n)
	} else {
		return 1 / cnt(x, -n)
	}
}
