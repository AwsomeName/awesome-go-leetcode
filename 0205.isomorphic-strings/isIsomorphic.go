package problem0205

func isIsomorphic(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	if m == 0 {
		return true
	}

	chars := make(map[byte]byte)
	chars2 := make(map[byte]byte)

	for i := 0; i < m; i++ {
		if x, ok := chars[s[i]]; ok && x != t[i] {
			return false
		}
		chars[s[i]] = t[i]
		if x, ok := chars2[t[i]]; ok && x != s[i] {
			return false
		}
		chars2[t[i]] = s[i]
	}

	return true
}
