package problem0015

import "sort"

func threeSum(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    Index := make(map[int]int, len(nums))

    for i := range nums {
        Index[nums[i]] = i
    }

    var last, tmp []int
    for i := 0; i < len(nums)-2; i++ {
        if nums[i] > 0 { return res }
        for i < len(nums)-3 && nums[i] == nums[i+1] && nums[i+1] == nums[i+2] { i++ }
        for j := i + 1; j < len(nums) - 1; j++{
            if nums[i] + nums[j] > 0 { break }
            if Index[0 - nums[i] - nums[j]] > j {
                tmp = []int{nums[i],nums[j]}
                if last == nil || tmp[0] != last[0] || tmp[1] != last[1] {
                    res = append(res, []int{nums[i], nums[j], nums[Index[0-nums[i]-nums[j]]]})
                    last = tmp
                }
            }
        }
    }

    return res
}
