package problem1006

func clumsy(N int) int {
    if N<3 {
        return N
    }

    ans := 0
    m1, m2 := N, N
    if N > 0 {
        m1 = N
        N--
    }else{
        return ans
    }

    if N > 0 {
        m2 = N * m1
        N--
    }else {
        ans += m1
        return ans
    }

    if N > 0 {
        m3 := m2/N
        ans += m3
        N--
    }else{
        ans += m2
        return ans
    }

    if N > 0 {
        ans += N
        N--
    }else{
        return ans
    }


    for N > 0 {
        fmt.Println(N, ans)
        n1,n2 := N, N

        if N > 0 {
            n1 = N
            N--
        }else{
            return ans
        }

        if N > 0 {
            n2 = N * n1
            N--
        }else {
            ans -= n1
            return ans
        }

        if N > 0 {
            n3 := n2/N
            ans -= n3
            N--
        }else{
            ans -= n2
            return ans
        }

        if N > 0 {
            ans += N
            N--
        }else{
            return ans
        }

        // if N > 0 {
        //     ans -= N
        //     N --
        // }else{
        //     return ans
        // }

    }



    return ans
}
