package problem0975

func oddEvenJumps(A []int) int {
    lenA := len(A)
    ans := 1;

    tmpMap := make(map[int]bool)
    tmpSort := make([]int, 0)
    for _,x := range A{
        if _,ok := tmpMap[x]; !ok {
            tmpSort = append(tmpSort, x)
            tmpMap[x] = true
        }
    }
    sort.Ints(tmpSort)

    higher := make([]bool, lenA)
    lower := make([]bool, lenA)

    higher[lenA - 1], lower[lenA - 1]  = true, true

    tMap := make(map[int]int)

    tMap[A[lenA-1]] = lenA - 1

    for i:= lenA - 2; i >= 0; i--{
        hi,ok  := Ceilling(tMap, A[i])
        if ok {
            higher[i] = lower[tMap[hi]]
        }
        lo,ok  := Floor(tMap, A[i])
        if ok {
            lower[i] = higher[tMap[lo]]
        }

        if higher[i] {
            ans ++
        }
        tMap[A[i]] = i
    }

    return ans
}

func find



func Ceilling (tm map[int]int, ti int) (int, bool){
    ans := math.MaxInt32
    for k,_ := range tm{
        if k >= ti {
            if k <= ans && k >= ti {
                ans = k
            }
        }
    }
    if ans < math.MaxInt32{
        return ans, true
    }
    return 0, false
}

func Floor (tm map[int]int, ti int) (int, bool){
    ans := 0
    for k,_ := range tm{
        if k <= ti {
            if k >= ans && k <= ti {
                ans = k
            }
        }
    }
    if ans > 0 {
        return ans, true
    }
    return 0, false
}

