// Package  problem1007

package problem1007

func minDominoRotations(A []int, B []int) int {
    ACnt := make([]int,7)
    BCnt := make([]int,7)

    for _,x := range A {
        ACnt[x]++
    }
    for _,x := range B {
        BCnt[x]++
    }
    index := 0
    for i:= 1; i<=6; i++{
        if ACnt[i] + BCnt[i] >= len(A){
            index = i
            break
        }
    }
    if index == 0 {
        return -1
    }

    ans := 0
    if ACnt[index] < BCnt[index] {
        flag := true
        for i,x := range B{
            if x != index{
                if A[i] != index{
                    flag = false
                    break
                }else{
                    ans ++
                }
            }
        }
        if flag {
            return ans
        }else{
            return -1
        }

    }else{
        flag := true
        for i,x := range A{
            if x != index{
                if B[i] != index{
                    flag = false
                    break
                }else{
                    ans ++
                }
            }
        }
        if flag {
            return ans
        }else{
            return -1
        }


    }

    return -1
}
