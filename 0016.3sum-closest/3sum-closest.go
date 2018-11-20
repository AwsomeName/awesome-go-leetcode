package problem0016

import "sort"
import "fmt"

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    result := nums[0] + nums[1] + nums[2]
    for i := range nums{
        if delta(result,target) == 0 { return result}
        for j := i + 1; j < len(nums)-1; j++{
            for x:=j+1; x <=len(nums)-1; x++{
                if delta(result, target) > delta( nums[i] + nums[j] + nums[x], target) {
                    result = nums[i]+ nums[j] + nums[x]
                    fmt.Println(result, nums[i], nums[j], nums[x],"-------------")
                }
                fmt.Println(result, nums[i], nums[j], nums[x])
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
