package problem0973

func kClosest(points [][]int, K int) [][]int {
	l := len(points)
	if l <= 0 || K > l+1 {
		return [][]int{}
	}
	dis := make([]int, l)

	for i, x := range points {
		dis[i] = x[0]*x[0] + x[1]*x[1]
	}

	sort.Ints(dis)
	ans := make([][]int, 0)
	d := dis[K-1]
	for _, x := range points {
		if K > 0 && x[0]*x[0]+x[1]*x[1] <= d {
			ans = append(ans, x)
			K--
		}
	}
	return ans
}
