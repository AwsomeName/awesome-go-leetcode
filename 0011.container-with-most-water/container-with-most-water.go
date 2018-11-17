package problem0011

func maxArea(height []int) int {

    max := 1
    tmp := 0
    for i := 0; i < len(height)-1; i++{
        for j := i + 1; j < len(height); j++ {
            tmp = min(height[i],height[j])
            tmp = tmp * (j - i)
            if tmp > max { max = tmp}
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
