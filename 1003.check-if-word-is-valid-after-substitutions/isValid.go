package problem1003

func isValid(S string) bool {
    A, B := 0, 0
    for _,x:= range S {
        if x == 'a' {
            A ++

        }else if x == 'b' {
            B ++
            if B > A {
                return false
            }

        }else if x== 'c'{
            A --
            B --
            if A < 0 || B < 0 {
                return false
            }
        }

    }

    if A > 0 || B > 0 {
        return false
    }

    return true
}
