package problem1000

var (
    gs []int
    cache map[int]int
    gk int
)

func dp(s, e int) int {
    i := (s<<8) | e
    if v, ok := cache[i]; ok { return v }
    l := e-s+1
    if l<gk { cache[i] = 0; return 0 }
    sum := 0
    for j:=s; j<=e; j++ { sum += gs[j] }
    d := l%(gk-1)
    if d == 0 { d=gk-1 }
    if d == 1 { d=gk }
    v := -1
    for l1:=1; l1<l; l1++ {
        t := l1 % (gk-1)
        if t==0 { t = gk-1 }
        if t>=d { continue }
        tv := dp(s, s+l1-1) + dp(s+l1, e)
        if v==-1 || v>tv { v = tv }
    }
    if d==gk { v += sum }
    cache[i] = v
    return v
}

func mergeStones(stones []int, K int) int {
    l := len(stones)
    if l==1 { return 0 }
    if l < K { return -1 }
    if ((l-1) % (K-1)) != 0 { return -1 }
    gs, gk = stones, K
    cache = make(map[int]int)
    return dp(0, l-1)
}

