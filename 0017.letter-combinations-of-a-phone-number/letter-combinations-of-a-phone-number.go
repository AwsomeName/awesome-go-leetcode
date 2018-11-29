package problem0017

var keyMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	ret := []string{""}

	for i := 0; i < len(digits); i++ {
		tmp := []string{}
		for j := 0; j < len(ret); j++ {
			for k := 0; k < len(keyMap[digits[i]]); k++ {
				tmp = append(tmp, ret[j]+keyMap[digits[i]][k])
			}
		}
		ret = tmp
	}
	return ret
}
