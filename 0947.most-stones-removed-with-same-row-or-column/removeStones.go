package problem0947

func removeStones(stones [][]int) int {
    if len(stones) == 1 { return 0 }
    move := 0

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
    for i:= range stones{
        tmp := false
        for j:= i+1; j < len(stones); j++{
            if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
                tmp = true
                break
            }
        }
        if tmp == true {
            move ++
        }
    }
    return move
}
