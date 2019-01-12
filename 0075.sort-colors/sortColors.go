package problem0075

func sortColors(nums []int) {
	cnt0, cnt1, cnt2 := 0, 0, 0
	for _, x := range nums {
		switch {
		case x == 0:
			cnt0++
		case x == 1:
			cnt1++
		case x == 2:
			cnt2++
		default:
			break
		}
	}
	for i := 0; i < cnt0; i++ {
		nums[i] = 0
	}
	for i := cnt0; i < cnt1+cnt0; i++ {
		nums[i] = 1
	}
	for i := cnt1 + cnt0; i < cnt1+cnt0+cnt2; i++ {
		nums[i] = 2
	}
}
