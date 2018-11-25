package problem0947

func removeStones(stones [][]int) int {
    if len(stones) == 1 { return 0 }
    move := 0
    var mark [1000]int
    for i:= range stones{
        for j:= i+1; j < len(stones); j++{
            if stones[i][0] > stones[j][0] {
                stones[i],stones[j] = stones[j],stones[i]
            }else if  stones[i][0] == stones[j][0] {
                if stones[i][1] > stones[j][1] {
                    stones[i],stones[j] = stones[j],stones[i]
                }
            }
        }
    }

    for i:= 0; i< len(stones) - 1; i++{
        tmp := false
        for j:= i+1; j < len(stones); j++{
            if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
                tmp = true
                if mark[i] ==0 && mark[j] == 0 {
                    move++
                    mark[i] = move
                    mark[j] = move
                }else if mark[i] == 0 {
                    mark[i] = mark[j]
                }else {
                    mark[j] = mark[i]
                }
            }
        }
        if tmp == false && mark[i] == 0{
            move ++
        }
    }
    return len(stones) - move
}
