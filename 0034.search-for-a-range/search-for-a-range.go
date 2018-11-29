package problem0034

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	left, right := 0, len(nums)-1
	est := -1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println("left:", left, "right:", right, "mid:", mid, "nums:", nums[mid])
		if nums[mid] == target {
			est = mid
			break
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		fmt.Println("after", "left:", left, "right:", right, "mid:", mid, "nums:", nums[mid])
	}
	fmt.Println("est:", est)
	left, right = est, est
	if est >= 0 {
		for left-1 >= 0 && nums[left] == nums[left-1] {
			left--
		}
		for right+1 <= len(nums)-1 && nums[right] == nums[right+1] {
			right++
		}
		return []int{left, right}
	}
	return []int{-1, -1}
}
