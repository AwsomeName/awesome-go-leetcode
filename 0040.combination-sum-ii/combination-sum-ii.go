package problem0040

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	solution := []int{}
	cs(candidates, solution, target, &res)

	return res
}

func cs(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		// target < candidates[0] 因为candidates是排序好的
		return
	}

	solution = solution[:len(solution):len(solution)]

	cs(candidates[1:], append(solution, candidates[0]), target-candidates[0], result)

	index := 0
	for index+1 < len(candidates) && candidates[index] == candidates[index+1] {
		index++
	}
	cs(candidates[index+1:], solution, target, result)
}
