package problem0011

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}

		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
