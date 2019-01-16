package problem0076

func minWindow(s string, t string) string {
	req := make(map[byte]int)
	hav := make(map[byte]int)
	total := len(t)
	for i := range t {
		req[t[i]]++
	}

	min := len(s) + 1
	res := ""
	for i, j, cnt := 0, 0, 0; j < len(s); j++ {
		if hav[s[j]] < req[s[j]] {
			cnt++
		}
		hav[s[j]]++
		// reduce i
		for i <= j && hav[s[i]] > req[s[i]] {
			hav[s[i]]--
			i++
		}
		tmp := j - i + 1
		//record window
		if cnt == total && tmp < min {
			min = tmp
			// idx = i
			res = s[i : j+1]
			// fmt.Println(idx,min)
		}
	}
	// if idx+min <= len(s){
	//     return s[idx:idx+min]
	// }else {return ""}
	return res
}
