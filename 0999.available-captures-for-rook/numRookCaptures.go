package problem0999

func numRookCaptures(board [][]byte) int {
    N := len(board)
    found := false
    X,Y := 0, 0
    for i:=0; i<N; i++ {
        for j:=0; j < N; j++{
            if board[i][j] == 'R'{
                X, Y = i, j
                found = true
                fmt.Println(X,Y)
                break
            }
        }
        if found {
            break
        }
    }
    cnt := 0
    x,y := X, Y
    for x >= 1 {
        x--
        if board[x][y] == 'p'{
            cnt++
            break
        }else if board[x][y] == 'B'{
            // cnt++
            break
        }else{
            continue
        }
    }
    x,y = X, Y
    for x < N-1 {
        x++
        if board[x][y] == 'p'{
            cnt++
            break
        }else if board[x][y] == 'B'{
            // cnt++
            break
        }else{
            continue
        }
    }
    x,y = X, Y
    for y >= 1 {
        y--
        if board[x][y] == 'p'{
            cnt++
            break
        }else if board[x][y] == 'B'{
            // cnt++
            break
        }else{
            continue
        }
    }
    x,y = X, Y
    for y < N-1 {
        y++
        if board[x][y] == 'p'{
            cnt++
            break
        }else if board[x][y] == 'B'{
            // cnt++
            break
        }else{
            continue
        }
    }
    return cnt
}
