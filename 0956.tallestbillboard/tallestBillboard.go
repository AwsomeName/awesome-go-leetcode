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
	//fmt.Println("rods:",rods)
	if len(rods) <= 2 && sum(rods)%2 == 1 {
		return 0
	}
	if len(rods) == 2 && rods[0] != rods[1] {
		return 0
	}
	target := sum(rods) / 2
	res := false
	find(rods, target, &res)
	//    if sum(rods)%2 ==0 && find(rods,target){
	if sum(rods)%2 == 0 && res {
		//fmt.Println("target:", target, "sum:", sum(rods))
		return target
	} else {
		//fmt.Println("else: target:", target, "sum:", sum(rods),"rods:",rods)
		for i := 0; i < len(rods); i++ {
			nRods := rods[1:]
			nRods = nRods[:len(nRods):len(nRods)]
			if i > 0 && i < len(nRods) {
				nRods[i], rods[0] = rods[0], nRods[i]
			}
			//fmt.Println("nRods:",nRods)
			sort.Ints(nRods)
			//            fmt.Println("call cnt:",nRods)
			r := cnt(nRods)
			if r > 0 {
				return r
			}
		}
		return 0
	}
}

func tallestBillboard(rods []int) int {
	fmt.Println("----------", rods, "---------")
	sort.Ints(rods)
	return cnt(rods)
}
