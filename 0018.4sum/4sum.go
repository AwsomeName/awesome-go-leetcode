package problem0018

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    var result [][]int

    for i:=0; i < len(nums) - 3 ;i++{
        if i>0 && nums[i] == nums[i-1] { continue }

        for j:=i+1; j < len(nums)-2; j++{
            if j > i+1 && nums[j] == nums[j-1] { continue }

            l,r := j+1, len(nums)-1
            for l < r {
                s := nums[i] + nums[j] + nums[l] + nums[r]
                switch {
                case s < target:
                    l++
                case s > target:
                    r--
                default:
                    result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
                    l, r = next(nums, l, r)
               }
            }
        }
    }

    return result
}

func next(nums []int, l, r int) (int, int) {
    for l < r {
        switch {
        case nums[l] == nums[l+1]:
            l++
        case nums[r] == nums[r-1]:
            r--
        default:
            l++
            r--
            return l,r
        }
    }
    return l,r
}
