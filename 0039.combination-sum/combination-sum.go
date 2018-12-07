package problem0039

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	solution := []int{}
	cs(candidates, solution, target, &res)
	return res
}

func cs(candidates []int, solution []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
		return
	}
	if len(candidates) == 0 || target < candidates[0] {
		return
	}
	solution = solution[:len(solution):len(solution)]
	cs(candidates, append(solution, candidates[0]), target-candidates[0], res)
	cs(candidates[1:], solution, target, res)
}
