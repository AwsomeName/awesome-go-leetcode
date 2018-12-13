package problem0049

var prime = []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}

func groupAnagrams(ss []string) [][]string {

	var encode func(s string) int

	encode = func(s string) int {
		c := 1
		for x := range s {
			c = c * prime[(s[x]-'a')]
		}
		return c
	}

	var ans [][]string

	tmp := make(map[int][]string)
	for _, s := range ss {
		c := encode(s)
		tmp[c] = append(tmp[c], s)
	}
	for _, s := range tmp {
		ans = append(ans, s)
	}

	return ans
}
