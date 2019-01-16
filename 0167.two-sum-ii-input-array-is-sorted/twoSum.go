package problem0167

func twoSum(numbers []int, target int) []int {
	n := numbers
	i, j := 0, len(n)-1
	for i < j && n[i]+n[j] != target {
		if n[i]+n[j] > target {
			j--
		} else {
			i++
		}
	}
	ans := []int{i + 1, j + 1}
	return ans
}
