package problem0005

func longestPalindrome(s string) string{
    if len(s) < 2 { return s}

    begin, maxLen := 0,1

    for i:=0; i<len(s);{
        if len(s)-i < maxLen/2 {
            break
        }
        b,e := i,i

        for e< len(s)-1 && s[e]==s[e+1]{
            e++
        }
        i = e + 1

        for e < len(s)-1 && b > 0 && s[e+1]==s[b-1]{
            e++
            b--
        }

        newLen := e +1 -b
        if newLen > maxLen{
            maxLen = newLen
            begin = b
        }

    }

    return s[begin:begin+maxLen]
}
