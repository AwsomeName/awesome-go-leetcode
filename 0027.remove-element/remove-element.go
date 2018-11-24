package problem0027

func removeElement(nums []int, val int) int {
    if len(nums)==0 { return 0 }
    i:=0
    j:=len(nums)-1
    for i < j {
        for i < j && nums[i] !=  val {
            i++
        }
        for j > i && nums[j] == val {
            j--
        }
        nums[i],nums[j] = nums[j], nums[i]
    }
    if j== len(nums)-1 && nums[j]!= val { i++ }
	return i
}
