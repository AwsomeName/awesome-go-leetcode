package problem0947

func removeStones(stones [][]int) int {
    if len(stones) == 1 { return 0 }
    move := 0
    var mark [1000]int
    for i := range stones{
        if mark[i] > 0 {continue}
        move ++
        var tmpQue [1000]int
        que_head := 0
        que_tail := 0
        tmpQue[que_head] = i
        for que_head <= que_tail {
            tmp := tmpQue[que_head]
            que_head++
            if mark[tmp] > 0 {continue}
            mark[tmp] = move
            for j:= range stones {
                if mark[j] > 0 { continue}
                if stones[tmp][0] == stones[j][0] || stones[tmp][1] == stones[j][1] {
                    que_tail++
                    tmpQue[que_tail]= j
                }
            }

        }
    }
    return len(stones) - move
}

