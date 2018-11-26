package problem0029

import "math"

func divide(dividend int, divisor int) int {
    if divisor  == 0 { return math.MaxInt32 }
    if dividend == 0 { return 0 }
    var res int
    var mSign,nSign,resSign int
    if dividend > 0 {mSign = 1} else if dividend < 0 { mSign = -1}
    if divisor  > 0 {nSign = 1} else if divisor  < 0 { nSign = -1}
    if mSign + nSign == 0 { resSign = -1 } else { resSign = 1}
    if dividend < 0 { dividend = -dividend }
    if divisor  < 0 { divisor  = -divisor  }
    res,_ = recursion_divide(dividend,divisor ,1)
    if res > math.MaxInt32 && resSign>0 { return math.MaxInt32}
    if res == math.MaxInt32 && resSign < 0 { return -res }
    if resSign > 0 { return res }else { return -res }
}

func recursion_divide(m,n,cnt int)(res, rmd int){
    switch{
    case m < n:
        return 0,m
    case n <=m && m < n+n:
        return cnt, m-n
    default:
        res, rmd = recursion_divide(m, n+n, cnt+cnt)
        if rmd >=  n {
            return res+cnt, rmd - n
        }
    }
    return
}

