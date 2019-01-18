package problem0093

func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				if check(s[:i+1]) && check(s[i+1:j+1]) && check(s[j+1:k+1]) && check(s[k+1:]) {
					tmp := s[:i+1] + "." + s[i+1:j+1] + "." + s[j+1:k+1] + "." + s[k+1:]
					ans = append(ans, tmp)
				}
			}
		}
	}
	return ans
}

func check(s string) bool {
	// n := len(s)
	// if n==0 || n >3 || (n==3 && s[0] > '2'){
	//     return false
	// }
	if len(s) > 4 || (len(s) > 1 && s[0] == '0') {
		return false
	}
	i, _ := strconv.Atoi(s)
	if i > 255 {
		return false
	}
	return true
}
