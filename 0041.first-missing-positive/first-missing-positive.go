package problem0041

func firstMissingPositive(nums []int) int {
	M := [1000]bool{}
	M[0] = true
	for _, x := range nums {
		if x > 0 {
			M[x] = true
		}
	}
	for i := 0; i < 1000; i++ {
		if M[i] == false {
			return i
		}
	}
	return 0
}
