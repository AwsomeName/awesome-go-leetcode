package problem0956

import "fmt"
import "sort"

func sum(rods []int) int {
	sum := 0
	for _, x := range rods {
		sum += x
	}
	return sum
}

func find(rods []int, target int, res *bool) {
	if *res || target == 0 {
		*res = true
		return
	}
	if len(rods) <= 0 || target < rods[0] {
		return
	}
	find(rods[1:], target-rods[0], res)
	find(rods[1:], target, res)
}

func cnt(rods []int) int {
	//fmt.Println("rods:",rods,"-------------------")
	if len(rods) < 2 || len(rods) == 2 && rods[0] != rods[1] {
		return 0
	}
	target := sum(rods) / 2
	res := false
	find(rods, target, &res)
	if sum(rods)%2 == 0 && res {
		//		fmt.Println("target:", target, "sum:", sum(rods))
		return target
	} else {
		//fmt.Println("else: target:", target, "sum:", sum(rods),"rods:",rods)
		res := 0
		for i := 0; i < len(rods); i++ {
			nRods := make([]int, len(rods)-1, len(rods)-1)
			for j := 0; j < len(rods)-1; j++ {
				if j >= i {
					nRods[j] = rods[j+1]
				} else {
					nRods[j] = rods[j]
				}
			}
			r := cnt(nRods)
			if len(nRods) == 10 {
				fmt.Println("nRods:", nRods)
			}
			if r > 0 && res < r {
				res = r
				if len(nRods) == 10 {
					fmt.Println("nRods:", nRods, res, rods, "i:", i)
				}
			}
		}
		return res
	}
}

func tallestBillboard(rods []int) int {
	fmt.Println("=================", rods, "=====================")
	sort.Ints(rods)
	return cnt(rods)
}
