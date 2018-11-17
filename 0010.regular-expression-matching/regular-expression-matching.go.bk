package problem0010

import "regexp"

func isMatch(s, p string) bool {
    reg,_ := regexp.Compile(p)
    index := reg.FindStringIndex(s)
    if index == nil {
        return false
    }
    if index[1] - index[0]  < len(s) {
        return false
    }
    return true
}
