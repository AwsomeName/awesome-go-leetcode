package problem0204

func countPrimes(n int) int {
	cnt := 0
	nP := make([]bool, n)
	// nP[2] = true
	for i := 2; i < n; i++ {
		if !nP[i] {
			cnt++
			for j := 2; i*j < n; j++ {
				nP[i*j] = true
			}
		}
	}
	return cnt
}
