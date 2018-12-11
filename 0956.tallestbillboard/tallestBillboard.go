package problem0956

import "fmt"
import "sort"
import "math"

func cnt(rods []int) int {
    var dp func(i,s int)int
    dp = func(i,s int) int {
        if i==len(rods) {
            if s==0 {
                return 0
            }else {
                return -math.MaxInt32
            }
        }
        maxAns := dp(i+1,s)
        tmp := dp(i+1,s-rods[i])
        if maxAns < tmp {
            maxAns = tmp
        }
        tmp = dp(i+1,s+rods[i]) + rods[i]
        if maxAns < tmp {
            maxAns = tmp
        }
        return maxAns
    }
    return dp(0,0)
}

func tallestBillboard(rods []int) int {
	fmt.Println("=================", rods, "=====================")
	sort.Ints(rods)
	return cnt(rods)
}
