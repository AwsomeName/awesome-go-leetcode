package problem0087

func isScramble(s1 string, s2 string) bool {
    if s1==s2 {return true}
    n := len(s1)

    if len(s2) != n { return false}

    count :=make([]int, 26)
    for i:=0; i < n; i++{
        count[int(s1[i]-'a')]++;
        count[int(s2[i]-'a')]--;
    }
    for _,x := range count{
        if x != 0 {
            return false
        }
    }

    for i:= 1; i< n; i++{
        if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:],s2[i:]) {
            return true
        }
        if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:],s2[:n-i]) {
            return true
        }
    }
    return false
}

