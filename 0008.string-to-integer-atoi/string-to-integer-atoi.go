package problem0008

import "math"
import "strings"

func myAtoi(str string) int {
	nonWhiteSpaceStr := strings.TrimSpace(str)
	if len(nonWhiteSpaceStr) == 0 {
		return 0
	}

	sign := 1
	if nonWhiteSpaceStr[0] == '-' {
		sign = -1
		nonWhiteSpaceStr = nonWhiteSpaceStr[1:]
	} else if nonWhiteSpaceStr[0] == '+' {
		nonWhiteSpaceStr = nonWhiteSpaceStr[1:]
	}

	if len(nonWhiteSpaceStr) == 0 {
		return 0
	}

	var cnvrt_int int

	if nonWhiteSpaceStr[0] > '9' || nonWhiteSpaceStr[0] < '0' {
		return 0
	}

	for i := range nonWhiteSpaceStr {
		if nonWhiteSpaceStr[i] >= '0' && nonWhiteSpaceStr[i] <= '9' {
			cnvrt_int = cnvrt_int*10 + (int(nonWhiteSpaceStr[i])-48)*sign
			if cnvrt_int > math.MaxInt32 {
				return math.MaxInt32
			} else if cnvrt_int < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			return cnvrt_int
		}
	}
	return cnvrt_int

}
