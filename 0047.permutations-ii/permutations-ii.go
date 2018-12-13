package problem0047

import "sort"

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	n := len(nums)
	vector := make([]int, n)
	taken := make([]bool, n)
	makePerUni(0, n, vector, nums, taken, &ans)
	return ans
}

func makePerUni(cur, n int, vector, nums []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}
	used := make(map[int]bool)
	for i := 0; i < n; i++ {
		if !taken[i] && !used[nums[i]] {
			taken[i] = true
			vector[cur] = nums[i]
			used[nums[i]] = true
			makePerUni(cur+1, n, vector, nums, taken, ans)
			taken[i] = false
		}
	}
}
