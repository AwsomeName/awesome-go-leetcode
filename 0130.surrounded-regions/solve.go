package problem0130

func solve(board [][]byte)  {
    m := len(board)
    if m == 0 {
        return
    }
    n := len(board[0])
    if n == 0 {
        return
    }


    marks := make([][]bool,m)
    for i:= range marks {
        marks[i] = make([]bool, n)
    }

    var check func(i,j int)
    check = func(i,j int) {
        if i > 0 {
            if board[i-1][j] == 'O' {
                if !marks[i-1][j] {
                    marks[i-1][j] = true
                    check(i-1,j)
                }
            }
        }
        if j > 0 {
            if board[i][j-1] == 'O' {
                if !marks[i][j-1] {
                    marks[i][j-1] = true
                    check(i,j-1)
                }
            }
        }
        if i < m-1 {
            if board[i+1][j] == 'O' {
                if !marks[i+1][j] {
                    marks[i+1][j] = true
                    check(i+1,j)
                }
            }
        }
        if j < n-1 {
            if board[i][j+1] == 'O' {
                if !marks[i][j+1] {
                    marks[i][j+1] = true
                    check(i,j+1)
                }
            }
        }

    }



    for i:=0; i<n; i++{
        if board[0][i] == 'O' && !marks[0][i] {
            marks[0][i] = true
            check(0,i)
        }
        if board[m-1][i] == 'O' && !marks[m-1][i] {
            marks[m-1][i] = true
            check(m-1,i)
        }
    }
    for j:=0; j<m; j++{
        if board[j][0] == 'O' && !marks[j][0] {
            marks[j][0] = true
            check(j,0)
        }
        if board[j][n-1] == 'O' && !marks[j][n-1] {
            marks[j][n-1] = true
            check(j,n-1)
        }
    }

    for i:=0; i<m; i++{
        for j := 0; j < n ; j++{
            if board[i][j] == 'O' && !marks[i][j] {
                board[i][j] = 'X'
            }
        }
    }
    return
}
