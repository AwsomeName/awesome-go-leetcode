package problem0974

func subarraysDivByK(A []int, K int) int {
	cnt := 0
	sMap := make(map[int]int)
	sMap[0] = 1
	sum := 0
	for _, a := range A {
		sum = (sum + a) % K
		if sum < 0 {
			sum += K
		}
		cnt += sMap[sum]
		sMap[sum]++
	}
	return cnt
}
