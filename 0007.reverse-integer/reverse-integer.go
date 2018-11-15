package problem0007

import "math"

func reverse(x int) int {
    if (x==0) { return 0 }
    y := 0
    for (x !=  0){
        y = y * 10 + x % 10
        x = x /10
    }
    if y > math.MaxInt32 || y < math.MinInt32 { return 0}
	return y
}
