package problem0090

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	i := 0
	for i < len(nums) {
		cnt := 0
		for cnt+i < len(nums) && nums[cnt+i] == nums[i] {
			cnt++
		}
		x := len(ans)
		for k := 0; k < x; k++ {
			tmp := make([]int, len(ans[k]))
			copy(tmp, ans[k])
			for j := 0; j < cnt; j++ {
				tmp = append(tmp, nums[i])
				ans = append(ans, tmp)
			}
		}
		i += cnt
	}
	return ans
}
