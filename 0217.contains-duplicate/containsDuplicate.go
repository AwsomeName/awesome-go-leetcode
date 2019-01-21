package problem0217

func containsDuplicate(nums []int) bool {
	key := make(map[int]bool)
	for _, x := range nums {
		if key[x] {
			return true
		}
		key[x] = true
	}
	return false
}
