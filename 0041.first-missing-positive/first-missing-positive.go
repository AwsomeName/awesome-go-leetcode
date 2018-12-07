package problem0041

func firstMissingPositive(nums []int) int {
	M := [1100]bool{}
	for _, x := range nums {
		if x > 0 {
			M[x] = true
		}
	}
	for i := 1; i < 1000; i++ {
		if M[i] == false {
			return i
		}
	}
	return 0
}
