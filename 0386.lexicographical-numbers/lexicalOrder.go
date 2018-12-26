package problem0386

func lexicalOrder(n int) []int {
	ans := make([]int, 0, n)

	lop(1, n, &ans)
	return ans
}

func lop(start, end int, ans *[]int) {
	for i := 0; i < 10; i++ {
		if len(*ans) >= end {
			return
		}
		if start+i <= end {
			*ans = append(*ans, start+i)
		}
		if (start+i)*10 <= end {
			lop((i+start)*10, end, ans)
		}
	}

}
