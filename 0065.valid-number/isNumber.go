package problem0065

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	s = selfTrimSpace(s)
	if len(s) == 0 {
		return false
	}
	idx := 0
	if s[idx] == '-' || s[idx] == '+' {
		idx++
	}
	if idx > len(s)-1 {
		return false
	}

	var A, Point, flot, E, afterE bool
	for idx < len(s) && s[idx] != '.' && s[idx] != 'e' {
		if s[idx] < '0' || s[idx] > '9' {
			return false
		}
		idx++
		A = true
	}
	// var Point part
	if idx < len(s) && s[idx] == '.' {
		Point = true
		idx++
	}
	// var flot part
	for idx < len(s) && s[idx] != 'e' {
		if s[idx] < '0' || s[idx] > '9' {
			// fmt.Println("s:",s, "idx:",idx, " :",string(s[idx]))
			// fmt.Println(A, Point, flot, E, afterE)
			return false
		}
		idx++
		flot = true
	}

	// var E part
	if idx < len(s) && s[idx] == 'e' {
		E = true
		idx++
	}
	if idx < len(s) && (s[idx] == '-' || s[idx] == '+') {
		idx++
	}

	// var afterE part
	for idx < len(s) {
		if s[idx] < '0' || s[idx] > '9' {
			return false
		}
		idx++
		afterE = true
	}

	if E && (!afterE || (!A && !flot)) {
		// fmt.Println(A, Point, flot, E, afterE)
		return false
	}

	if Point && (!A && !flot && !E) {
		// fmt.Println(A, Point, flot, E, afterE)
		return false
	}

	return true
}

func selfTrimSpace(s string) string {
	ss := make([]byte, 0)
	idx, idx2 := 0, len(s)-1
	for idx < len(s) && s[idx] == ' ' {
		idx++
	}
	for s[idx2] == ' ' && idx2 >= idx {
		idx2--
	}
	for i := idx; i <= idx2; i++ {
		ss = append(ss, s[i])
	}
	return string(ss)
}
