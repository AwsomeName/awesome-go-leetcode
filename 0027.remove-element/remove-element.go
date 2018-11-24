package problem0027

func removeElement(nums []int, val int) int {
    i:=0
    j:=len(nums)-1
    for i < j -1 {
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
