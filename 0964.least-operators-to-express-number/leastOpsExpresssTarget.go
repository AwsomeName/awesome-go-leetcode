package problem0964

func leastOpsExpressTarget(x int,target int) int{
    r1 := make([]int,0)
    for target > 0 {
        r1 = append(r1, target%x)
        target /= x
    }
    n := len(r1)
    pos := r1[0] *2
    neg := (x-r1[0])*2

    for i:=1; i<n; i++{
        pos,neg = min(r1[i]*i + pos, r1[i]*i +i+neg), min((x-r1[i])*i+pos,(x-r1[i])*i+i+neg-2*i)
    }
    return min(pos,n+neg)-1
}

func min(a,b int)int{
    if a < b {
        return a
    }else {
        return b
    }
}
