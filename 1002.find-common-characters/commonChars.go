package problem1002

func commonChars(A []string) []string {
    n := len(A)
    cnt := make([]map[byte]int,n)
    for i:= 0;i<n;i++{
        cnt[i] = make(map[byte]int)
    }
    for i:= range A{
        for j:= range A[i]{
            cnt[i][A[i][j]]++
        }
    }

    ans := make([]string,0)
    for k,v := range cnt[0] {
        if v > 0 {
            tmp := v
            flag := true
            for i:= 1; i< n ; i++{
                vv := cnt[i][k]
                if vv <= 0 {
                    flag = false
                    break
                }
                if vv < tmp {
                    tmp = vv
                }
            }
            if flag {
                for i := 0 ; i < tmp; i++{
                    ans = append(ans, string(k))
                }
            }
        }

    }

    return ans
}
