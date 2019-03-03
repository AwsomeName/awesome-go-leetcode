package problem0997

func findJudge(N int, trust [][]int) int {
    trustOthers := make([]int, N)
    toBeTrust := make([]int, N)
    for _,x := range(trust){
        trustOthers[x[0]-1]++
        toBeTrust[x[1]-1]++
    }
    // fmt.Println(trustOthers)
    // fmt.Println(toBeTrust)
    for i:=0; i<N; i++{
        if trustOthers[i] == 0 && toBeTrust[i] == N-1 {
            return i+1
        }
    }
    return -1

}
