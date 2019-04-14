package problem1031

func numEnclaves(A [][]int) int {
    ans := 0
    n := len(A)
    if n == 0 {
        return 0
    }
    m := len(A[0])
    if m == 0 {
        return 0
    }

    marks := make([][]bool, n)
    checked := make([][]bool, n)
    for i:= range marks{
        marks[i] = make([]bool, m)
        checked[i] = make([]bool, m)
    }

    for i:=0; i < n; i++{
        for j := 0 ; j < m; j ++ {
            if A[i][j] == 0 {
                marks[i][j] = true
                continue
            }else {
                if marks[i][j] || checked[i][j]{
                    continue
                // }else if checked[i][j] {
                //     continue
                }else{
                    tmp, flag := check(i,j, A)
                    fmt.Println(i,j)
                    fmt.Println(tmp)
                    fmt.Println(flag, "------------------------------")
                    for k := range tmp {
                        fmt.Println(A[tmp[k][0]][tmp[k][1]])
                        marks[tmp[k][0]][tmp[k][1]] = flag
                        checked[tmp[k][0]][tmp[k][1]] = true
                    }
                }
            }
        }
    }

    fmt.Println(marks)
    fmt.Println(checked)
    fmt.Println(A)

    for i := 0; i < n ; i++{
        for j :=0; j < m; j++{
            if !marks[i][j] {
                ans ++
            }
        }
    }

    return ans
}

func check(i, j int, A [][]int) ([][]int, bool) {
    n, m  := len(A), len(A[0])
    ans := make([][]int, 0)
    flag := false

    checked := 0
    record := make(map[string]bool)
    tmp := toString(i,j)
    record[tmp] = true
    ans = append(ans, []int{i,j})

    for checked < len(ans){
        i,j = ans[checked][0],  ans[checked][1]

        if iftrue(i,j, n, m) {
            flag = true
        }

        // check i-1
        if i-1 >=0 {
            s := i-1
            ss := toString(s, j)
            if !record[ss] && A[s][j] == 1{
                record[ss] = true
                ans = append(ans, []int{s,j})
            }
        }

        // check i+1
        if i+1 < n {
            s := i+1
            ss := toString(s, j)
            if !record[ss] && A[s][j] == 1 {
                record[ss] = true
                ans = append(ans, []int{s,j})
            }
        }

        // check j-1
        if j-1 >=0 {
            s := j-1
            ss := toString(i, s)
            if !record[ss] && A[i][s] == 1{
                record[ss] = true
                ans = append(ans, []int{i,s})
            }
        }

        // check j+1
        if j+1 < m {
            s := j+1
            ss := toString(i, s)
            if !record[ss] && A[i][s] == 1{
                record[ss] = true
                ans = append(ans, []int{i,s})
            }
        }

        checked++
    }



    return ans, flag
}

func toString(i, j int) string{
    return string(i) + "," + string(j)
}

func iftrue(i,j, n, m int) bool {
    if i==0 || i == n-1 || j == 0 || j == m-1 {
        return true
    }
    return false
}
