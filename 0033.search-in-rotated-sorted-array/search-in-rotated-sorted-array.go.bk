package problem0033

import "fmt"

func search(nums []int, target int) int {
    if len(nums) < 1 { return -1}
    if len(nums)==1 {
       if nums[0] == target { return 0} else { return -1}
    }
    tail := nums[len(nums)-1]
    head := nums[0]
    if target < head && target > tail { return -1 }
    res := 0
    left,right := 0, len(nums)-1
    for left < right{
        res = (left + right) /2
        fmt.Println("before> left:",left,"right:",right,"res:",res,"target:",target,"nums[res]:",nums[res])
        if nums[res] == target { return res }
        if nums[res] < target && nums[res] > head || nums[res] > target && nums[res] < tail || nums[res] > tail && target < tail{
            left = res + 1
        }else {
            right = res
        }
        fmt.Println("after> left:",left,"right:",right,"res:",res,"target:",target,"nums[res]:",nums[res])
    }
    if res > 1 && nums[res-1] == target { return res-1 }
    return -1
}
