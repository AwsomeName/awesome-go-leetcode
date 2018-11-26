package problem0029

import "math"

func divide(m, n int) int {
    if n == 0 { return math.MaxInt32 }
    if m == 0 { return 0 }
    var res int
    var mSign,nSign,resSign int
    if m > 0 {mSign = 1} else if m < 0 { mSign = -1}
    if n > 0 {nSign = 1} else if n < 0 { nSign = -1}
    if mSign + nSign == 0 { resSign = -1 } else { resSign = 1}
    if m < 0 { m = -m }
    if n < 0 { n = -n }
    res,_ = recursion_divide(m,n,1)
    if res > math.MaxInt32 { return math.MaxInt32}
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

