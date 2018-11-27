package problem0032
import "fmt"

func longestValidParentheses(s string) int {
    if  s=="" { return 0 }
    res, tmpCnt, leftCnt := 0,0,0
    head := 0
    for i,char := range s{
        if char=='(' {
            if i>1 && s[i-1]==')'{
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
        fmt.Println(char)
        fmt.Println("head:",head,"tmpCnt:",tmpCnt,"leftCnt:",leftCnt,"res:",res)
    }
    fmt.Println("head:",head,"tmpCnt:",tmpCnt,"leftCnt:",leftCnt,"res:",res)
    if s[len(s)-1]!='(' {
        res -= leftCnt
    }
    return res*2
}
