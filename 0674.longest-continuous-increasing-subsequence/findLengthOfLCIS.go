package problem0674

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	ans := 1
	cnt := 1
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			cnt++
			if ans < cnt {
				ans = cnt
			}
		} else {
			cnt = 1
		}

	}

	return ans
}
