package problem0189

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reserve(nums)
	reserve(nums[:k])
	reserve(nums[k:])
}

func reserve(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		// fmt.Println(i,n-i-1)
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
