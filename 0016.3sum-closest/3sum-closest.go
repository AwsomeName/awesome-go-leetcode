package problem0016

import "sort"

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    result := nums[0] + nums[1] + nums[2]
    for i := range nums{
        if nums[i] >= target { break }
        for j := i + 1; j < len(nums)-1; j++{
            l := j + 1
            r := len(nums) - 1
            if l > r { break }
            for x:=l; x <=r; x++{
                if delta(result, target) > delta( nums[i] + nums[j] + nums[x], target) {
                    result = nums[i]+ nums[j] + nums[x]
                }
            }
            for j < len(nums) && nums[j-1] == nums[j] { j++ }
        }
        for i > 0 && i < len(nums) && nums[i-1] == nums[i] { i++ }
    }
    return result
}

func delta(a int, b int) int{
    if a > b {
        return a - b
    } else {
        return b - a
    }
}
