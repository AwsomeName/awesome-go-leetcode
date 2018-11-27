package problem0031

import "math"

func nextPermutation(nums []int) {
    index := -1
    for i:= 0; i< len(nums)-1; i++{
        if nums[i] < nums[i+1] {
            index = i
        }
    }
    r_index := 0
    if index >= 0 {
        tmp_max := math.MaxInt32
        for i:= index+1; i < len(nums); i++{
            if nums[i] > nums[index] && nums[i] < tmp_max {
                tmp_max = nums[i]
                r_index = i
            }
        }
        nums[index],nums[r_index] =  nums[r_index],nums[index]
        reverseArray(nums[index+1:len(nums)])

        return
    }
    reverseArray(nums)
    return
}

func reverseArray(a []int){
    for i,j:=0,len(a)-1; i < j; {
        a[i],a[j] = a[j],a[i]
        i++
        j--
    }
}
