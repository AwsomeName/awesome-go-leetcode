package problem0190

func reverseBits(num uint32) uint32 {
	var h uint32
	h = 1
	h = h << 31
	var ans uint32

	for i := 0; i < 32; i++ {
		if h&num == 0 {
			ans = ans >> 1
		} else {
			ans = ans >> 1
			ans = ans | h
		}
		num = num << 1
	}
	return ans
}
