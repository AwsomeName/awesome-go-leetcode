package problem0058

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	cnt := 0
	for i := len(s) - 1; i >= 0 && s[i] == ' '; i-- {
		cnt++
	}
	for i := len(s) - 1 - cnt; i >= 0; i-- {
		if s[i] == ' ' {
			return len(s) - i - 1 - cnt
		}
	}
	return len(s) - cnt
}
