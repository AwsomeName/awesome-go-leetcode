package problem0032

func longestValidParentheses(s string) int {
    res, tmpCnt, leftCnt := 0,0,0
    head := 0
    for _,char := range s{
        if char=='(' {
            if head==0 {
                leftCnt = tmpCnt
            }
            head ++
        } else {
            if head > 0 {
                head --
                tmpCnt++
            }else {
                tmpCnt = 0
            }
            if head == 0 {
                leftCnt = 0
            }
        }
        if tmpCnt > res {
            res = tmpCnt
        }
    }
    if s[len(s)-1]==')' {
        res -= leftCnt
    }
    return res*2
}
