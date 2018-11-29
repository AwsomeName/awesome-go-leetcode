package problem0033

import "fmt"

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	lenN := len(nums)
	left, right := 0, lenN-1

	for left < right {
		midm := (left + right) / 2
		if nums[midm] > nums[right] {
			left = midm + 1
		} else {
			right = midm
		}
	}
	fmt.Println("left:", left, "nums[min]:", nums[left])
	rot_idx := left
	left, right = 0, lenN-1
	for left <= right {
		mid := (left + right) / 2
		rot_mid := (mid + rot_idx) % lenN
		fmt.Println("rot_mid:", rot_mid, "mid:", mid, "left:", left, "right:", right, "nums[rot_mid]:", nums[rot_mid])
		if nums[rot_mid] == target {
			return rot_mid
		}
		if nums[rot_mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
