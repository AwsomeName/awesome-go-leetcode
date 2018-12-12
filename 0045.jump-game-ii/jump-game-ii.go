package problem0045

import "math"

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 3 {
		return len(nums) - 1
	}
	jumps := make([]int, len(nums))
	jumps[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		jumps[i] = math.MaxInt32
		for j := i + 1; j < len(nums) && j <= i+nums[i]; j++ {
			tmp := jumps[j] + 1
			if jumps[i] > tmp {
				jumps[i] = tmp
			}
		}
	}
	return jumps[0]
}
