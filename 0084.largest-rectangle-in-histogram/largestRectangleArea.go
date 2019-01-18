package problem0084

// func largestRectangleArea(heights []int) int {
//     // Areas := make(map[int]int)
//     maxArea := 0
//     for i := 0; i < len(heights); i++{
//         m, n := i, i
//         for m-1 >=0 && heights[m-1] >= heights[i]{
//             m--
//         }
//         for n+1 < len(heights) && heights[n+1] >= heights[i]{
//             n++
//         }
//         width := n-m + 1
//         area := width * heights[i]
//         // Areas[heights[i]] = max(Areas[heights[i]], area)
//         maxArea = max(maxArea, area)
//     }
//     return maxArea
// }

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	L, R, stk := make([]int, len(heights)), make([]int, len(heights)), []int(nil)
	for i := 0; i < len(heights); i += 1 {
		h := heights[i]
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= h {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			L[i] = 0
		} else {
			L[i] = stk[len(stk)-1] + 1
		}
		if len(stk) == 0 || heights[stk[len(stk)-1]] != h {
			stk = append(stk, i)
		}
	}
	stk = []int(nil)
	for i := len(heights) - 1; i >= 0; i -= 1 {
		h := heights[i]
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= h {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			R[i] = len(heights) - 1
		} else {
			R[i] = stk[len(stk)-1] - 1
		}
		if len(stk) == 0 || heights[stk[len(stk)-1]] != h {
			stk = append(stk, i)
		}
	}
	ans := 0
	for i := 0; i < len(heights); i += 1 {
		ans = Max(ans, heights[i]*(R[i]-L[i]+1))
	}
	return ans
}
