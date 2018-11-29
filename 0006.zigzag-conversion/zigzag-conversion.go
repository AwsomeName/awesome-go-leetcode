package problem0006

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var buffer bytes.Buffer
	var j, direct int
	var superstr [1000]string
	for i := 0; i < len(s); i++ {
		if j == 0 {
			direct = 1
		} else if j == numRows-1 {
			direct = -1
		}
		superstr[j] += string(s[i])
		j = j + direct
	}
	for i := 0; i < numRows; i++ {
		buffer.WriteString(superstr[i])
	}
	return buffer.String()
}
