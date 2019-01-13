package problem0078

func subsets(nums []int) [][]int {
	var n, i, j, nlen uint
	nlen = uint(len(nums))
	n = 1 << nlen
	ans := make([][]int, 0)
	for i = 0; i < n; i++ {
		tmp := make([]int, 0)
		for j = 0; j < nlen; j++ {
			if i&(1<<j) > 0 {
				tmp = append(tmp, nums[j])
			}
		}
		ans = append(ans, tmp)
	}
	// tmp := []int{}
	// ans = append(ans,tmp)
	return ans
}
