package problem0014

func longestCommonPrefix(strs []string) string {
	res := ""
	j := 0
	if len(strs) == 0 {
		return ""
	}
	for {
		for i := 0; i < len(strs); i++ {
			if j >= len(strs[i]) {
				return res
			}
		}
		for i := 1; i < len(strs); i++ {
			if strs[i][j] != strs[0][j] {
				return res
			}
		}
		res += string(strs[0][j])
		j++
	}
	return res
}
