package problem0034

func searchRange(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    est:= -1
    for left < right {
        mid:= (left + right)/2
        if nums[mid] == target {
            est = mid
            break
        }
        if nums[mid] < target {
            left =mid +1
        }else {
            right = mid - 1
        }
    }
     left,right = est,est
    if est >=0 {
        for left-1>=0 && nums[left]==nums[left-1] {
            left--
        }
        for right+1<len(nums)-1 && nums[right] == nums[right+1] {
            right++
        }
        return []int{left,right}
    }
     return []int{-1,-1}
}
