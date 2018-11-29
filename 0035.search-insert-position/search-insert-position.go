package problem0035

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		{
			return 0
		}
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println("left:", left, "right:", right, "mid:", mid, "nums:", nums[mid], "target:", target)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			if mid == len(nums)-1 {
				return mid + 1
			}
			if mid < len(nums)-1 && nums[mid+1] > target {
				return mid + 1
			}
			left = mid + 1
		} else {
			if mid == 0 {
				return 0
			}
			if mid > 0 && nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		}
		fmt.Println("-----after-----:", "left:", left, "right:", right, "mid:", mid, "nums:", nums[mid], "target:", target)
	}
	return left
}
