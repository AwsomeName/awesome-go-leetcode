package problem1004

func longestOnes(A []int, K int) int {
    B := make([]int,0)
    cnt := 0
    for i:= range A{
        if A[i]==0 {
            if cnt > 0 {
                B = append(B, cnt)
                cnt = 0
            }
            B = append(B , 0)
        }else{
            cnt ++
            if i==len(A)-1 {
                B = append(B, cnt)
            }
        }
    }
    fmt.Println(B)

    C := make([]int, len(B))
    for i:= 0; i< len(C); i++{
        k := K
        tmp := 0
        j := i
        for k >= 0 && j < len(B) {
            if B[j] == 0 {
                k--
                j++
                if k >= 0 {
                    tmp ++
                }
                continue
            }else {
                tmp += B[j]
                j++
            }
        }
        C[i] = tmp
    }

    ans := 0
    for _,x := range C {
        if ans < x {
            ans = x
        }
    }
    return ans
}
