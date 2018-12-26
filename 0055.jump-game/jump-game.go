package problem0055

func canJump(nums []int) bool {
	marks := make([]bool, len(nums))
	marks[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i; j <= i+nums[i] && j < len(nums); j++ {
			if marks[j] {
				marks[i] = true
				break
			}
		}
	}
	return marks[0]
}
