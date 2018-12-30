package problem0966

func spellchecker(wordlist []string, queries []string) []string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	ans := make([]string, len(queries))
	WLMap := make(map[string][]int)
	LowerMap := make(map[string][]int)
	VowMap := make(map[string][]int)

	for i, x := range wordlist {
		WLMap[x] = append(WLMap[x], i)
		l := strings.ToLower(x)
		LowerMap[l] = append(LowerMap[l], i)
		var ll string
		for j := range l {
			if vowels[l[j]] {
				ll += "#"
			} else {
				ll += string(l[j])
			}
		}

		VowMap[ll] = append(VowMap[ll], i)
	}

	// vowels := map[byte]bool{'a':true,'e':true,'i':true,'o':true,'u':true,'A':true,'E':true,'I':true,'O':true,'U':true}

	for i, x := range queries {
		if WLMap[x] != nil {
			ans[i] = wordlist[WLMap[x][0]]
			continue
		}

		l := strings.ToLower(x)
		if LowerMap[l] != nil {
			ans[i] = wordlist[LowerMap[l][0]]
			continue
		}

		var ll string
		for j := range l {
			if vowels[l[j]] {
				ll += "#"
			} else {
				ll += string(l[j])
			}
		}

		if VowMap[ll] != nil {
			ans[i] = wordlist[VowMap[ll][0]]
			continue
		}

	}

	return ans
}
