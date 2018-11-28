package problem0032

import "fmt"

func longestValidParentheses(s string) int {
    if  s=="" { return 0 }
    var res int
    lParStk := make([]int, len(s))
    lParHdr := 0
    marks := make([]bool, len(s))
    for i,par := range s{
        if par == '(' {
            lParStk[lParHdr] = i
            lParHdr++
        }else{
            if lParHdr > 0 {
                lParHdr--
                marks[lParStk[lParHdr]],marks[i] = true, true
            }
        }
    }
    fmt.Println("match:",s)
    fmt.Println("match:",marks)
    for i,mark := range marks{
        tmpCnt:=0
        for mark && i<len(s) {
            tmpCnt++
            i++
            if i<len(s) {
                mark = marks[i]
            }
        }
        if tmpCnt >  res {
            res = tmpCnt
        }
    }
    return res
}
