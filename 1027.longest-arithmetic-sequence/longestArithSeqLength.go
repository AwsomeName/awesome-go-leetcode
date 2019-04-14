package problem1027

func longestArithSeqLength(A []int) int {
    if len(A) == 0 {
        return 0
    }

    // ans := make([]int, 10)
    amap := make(map[int]map[int]int)//存差值的最

    for _,x := range A{
        if _,ok := amap[x]; !ok {
            amap[x] = make(map[int]int)
        }

        for k,_ := range amap {// 对于每个值，还要考察它的差值
                cha := k - x
                if _,ok := amap[x][cha]; !ok {
                    amap[x][cha] = 2
                }
                if ttt,ok := amap[k][cha]; ok {
                    if ttt >= amap[x][cha] && cha != 0 {
                        amap[x][cha] = ttt + 1
                    }
                }
        }
    }
    // fmt.Println(amap)

    return max(amap)
}

func max(ans map[int]map[int]int) int{
    a := 0
    for _,x := range ans {
        for _,y := range x{
            if a < y {
                a = y
            }
        }
    }
    return a
}
