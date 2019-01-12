package problem0071

func simplifyPath(path string) string {
	stack := make([]string, 0)
	idx := 0
	var tmp string
	for _, x := range path {
		if x == '/' {
			if tmp == ".." {
				idx--
				if idx < 0 {
					idx = 0
				}
			} else if tmp != "." && len(tmp) > 0 {
				tmp = "/" + tmp
				if len(stack) < idx+1 {
					stack = append(stack, tmp)
				} else {
					stack[idx] = tmp
				}
				idx++
			}
			tmp = ""
			continue
		}
		tmp = tmp + string(x)
	}
	if tmp == ".." {
		idx--
		if idx < 0 {
			idx = 0
		}
	} else if tmp != "." && len(tmp) > 0 {
		tmp = "/" + tmp
		if len(stack) < idx+1 {
			stack = append(stack, tmp)
		} else {
			stack[idx] = tmp
		}
		idx++
	}
	var ans string

	for i := 0; i < idx; i++ {
		ans += stack[i]
	}
	if len(ans) <= 0 {
		ans = "/"
	}
	return ans
}
