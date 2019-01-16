package problem0169

func majorityElement(nums []int) int {
	n := len(nums)
	dock := make(map[int]int)
	for _, x := range nums {
		dock[x]++
		if dock[x] > n/2 {
			return x
		}
	}
	return 0
}
