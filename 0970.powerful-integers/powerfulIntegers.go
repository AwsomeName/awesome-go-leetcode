package problem0970

func powerfulIntegers(x int, y int, bound int) []int {
	if x == 1 && y == 1 {
		if bound >= 2 {
			return []int{2}
		} else {
			return []int{}
		}
	}
	ans := make([]int, 0)
	mRes := make(map[int]bool)

	var gen func(i, j int)

	gen = func(i, j int) {
		tmp := f(x, i) + f(y, j)
		// fmt.Println(x,i,y,j,tmp,f(x,i),f(y,j))
		if tmp <= bound && !mRes[tmp] {
			ans = append(ans, tmp)
			mRes[tmp] = true
		}
		if tmp > bound {
			return
		}
		if x > 1 {
			gen(i+1, j)
		}
		if y > 1 {
			gen(i, j+1)
		}
	}
	gen(0, 0)
	return ans
}

func f(a, b int) int {
	//return a^b
	if b == 0 || a == 1 {
		return 1
	}
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = res * a
		}
		a = a * a
		b /= 2
	}
	return res
}
