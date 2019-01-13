package problem0097

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	var check func(s1 string, s2 string, s3 string) bool
	check = func(s1 string, s2 string, s3 string) bool {
		if (len(s1) == 0 && s2 == s3) || (len(s2) == 0 && s1 == s3) || (len(s1) == 0 && len(s2) == 0 && len(s3) == 0) {
			return true
		}
		if (len(s1) == 0 && s2 != s3) || (len(s2) == 0 && s1 != s3) {
			return false
		}
		if s1[0] == s3[0] && s2[0] != s3[0] {
			return check(string(s1[1:]), s2, string(s3[1:]))
		} else if s2[0] == s3[0] && s1[0] != s3[0] {
			return check(string(s2[1:]), s1, string(s3[1:]))
		} else if s1[0] == s3[0] && s2[0] == s3[0] {
			return check(string(s1[1:]), s2, string(s3[1:])) || check(string(s2[1:]), s1, string(s3[1:]))
		} else {
			return false
		}
	}

	return check(s1, s2, s3)
}
