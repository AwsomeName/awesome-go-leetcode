package problem0219

func containsNearbyDuplicate(nums []int, k int) bool {
	idx := make(map[int]int)
	for i, x := range nums {
		if ii, ok := idx[x]; ok {
			if i-ii <= k {
				fmt.Println(i, ii, nums)
				return true
			}
		}
		idx[x] = i
	}
	return false
}
