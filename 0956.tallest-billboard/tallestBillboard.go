package problem0956

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func help(rods []int) map[int]int {

	dp := make(map[int]int)
	dp[0] = 0
	for _, x := range rods {
		cur := make(map[int]int)
		for d, y := range dp {
			cur[d] = y
		}
		//fmt.Println("before dp:",dp)
		for d, y := range cur {
			//fmt.Println("d,y,x:",d,y,x)
			dp[d+x] = max(dp[x+d], y)
			dp[abs(d-x)] = max(dp[abs(d-x)], y+min(d, x))
		}
		//fmt.Println("after dp:",dp)
	}
	return dp
}

func tallestBillboard(rods []int) int {
	fmt.Println("=================", rods, "=====================")
	res := 0
	dp := help(rods[len(rods)/2:])
	dp2 := help(rods[:len(rods)/2])
	for i, x := range dp {
		if y, ok := dp2[i]; ok {
			tmp := i + x + y
			if res < tmp {
				res = tmp
			}
		}
	}
	return res
}
