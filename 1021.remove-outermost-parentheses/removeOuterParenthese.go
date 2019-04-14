package problem1021

func removeOuterParentheses(S string) string {
    var ans string

    if len(S) <= 1 {
        return ans
    }

    left := 0
    // right := 1
    for i := 0; i < len(S); i++ {
        if S[i] == '(' {
            left ++
            j := i
            for left > 0 && j < len(S){
                j++
                if S[j] == '(' {
                    left++
                }else if S[j] == ')' {
                    left--
                }
                if left ==0 && j +1 < len(S) && S[j+1] == ')' {
                    left++
                }
            }
            tmp := S[i+1:j]
            fmt.Println(tmp)
            ans += string(tmp)
            i=j

        }else if S[i] == ')'{
            continue
        }

    }

    return ans

}
