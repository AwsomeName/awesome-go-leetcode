package problem0068

func fullJustify(words []string, maxWidth int) []string {
	if len(words) <= 0 {
		return []string{}
	}

	ans := make([]string, 0)
	idx := 0

	for idx < len(words) {
		var tmp string
		var last bool

		tmpIdx := idx
		tmpCnt := len(words[idx])
		for tmpIdx+1 < len(words) && tmpCnt+1+len(words[tmpIdx+1]) <= maxWidth {
			tmpIdx++
			tmpCnt += (len(words[tmpIdx]) + 1)
		}
		if tmpCnt <= maxWidth && tmpIdx == len(words)-1 {
			last = true
		}
		Cnt := tmpIdx - idx
		SpaceCnt := 0
		baseSpace := 1
		extSpace := 0
		if Cnt >= 1 {
			SpaceCnt = maxWidth - (tmpCnt - (Cnt))
			baseSpace = SpaceCnt / Cnt
			extSpace = SpaceCnt % Cnt
		}
		fmt.Println(Cnt, SpaceCnt, baseSpace, extSpace)
		for i := idx; !last && i <= tmpIdx; i++ {
			tmp = tmp + words[i]
			if i < tmpIdx {
				for x := 0; x < baseSpace; x++ {
					tmp = tmp + " "
				}
				if extSpace > 0 {
					extSpace--
					tmp = tmp + " "
				}
			}
		}

		for i := idx; last && i <= tmpIdx; i++ {
			tmp = tmp + words[i]
			if i < tmpIdx {
				tmp = tmp + " "
			}
		}

		for x := len(tmp); x < maxWidth; x++ {
			tmp = tmp + " "
		}

		idx = tmpIdx + 1
		ans = append(ans, tmp)
	}
	return ans
}
