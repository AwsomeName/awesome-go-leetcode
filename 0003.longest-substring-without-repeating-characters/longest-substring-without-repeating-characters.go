package problem0003

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	max_cnt := 0
	header_index := 0
	for i := 0; i < 255; i++ {
		location[i] = -1
	}

	tmp_index := 0
	for i := 0; i < len(s); i++ {
		tmp_index = location[s[i]]
		if tmp_index >= header_index {
			header_index = tmp_index + 1
		} else {

			if i-header_index+1 > max_cnt {
				max_cnt = i - header_index + 1
			}
		}

		location[s[i]] = i
	}
	return max_cnt
}
