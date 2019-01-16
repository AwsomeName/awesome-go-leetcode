package problem

func singleNumber(nums []int) int {
	var result int
	for _, val := range nums {
		result ^= val
	}
	return result
}
