package problem0081

func search(nums []int, target int) bool {
	if len(nums) < 1 {
		return false
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		} else {
			return false
		}
	}
	lenN := len(nums)
	left, right := 1, lenN-2

	for left < right && nums[left] >= nums[left-1] && nums[right] <= nums[right+1] {
		left++
		right--
	}
	rot_idx := 0
	// rot_idx:=left

	fmt.Println(left, right)
	if nums[left] < nums[left-1] {
		rot_idx = left
	} else if nums[right] > nums[right+1] {
		rot_idx = right + 1
	}

	left, right = 0, lenN-1
	for left <= right {
		mid := (left + right) / 2
		rot_mid := (mid + rot_idx) % lenN
		if nums[rot_mid] == target {
			return true
		}
		if nums[rot_mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false

}
