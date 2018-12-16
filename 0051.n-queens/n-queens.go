package problem0051

import "fmt"

func solveNQueens(n int) [][]string {
    if n==0 {
        return [][]string{}
    }
    var res [][]string
    cols := make([]bool,n)
    d1 := make([]bool,2*n)
    d2 := make([]bool,2*n)
    board := make([]string,n)
    dfs(0,n, cols, d1, d2, board, &res)
    return res
}

func dfs(idx,n int, cols,d1,d2 []bool, board []string, res *[][]string){
    if idx == n {
        tmp := make([]string,n)
        copy(tmp,board)
        *res = append(*res, tmp)
        fmt.Println("res:",*res)
        fmt.Println("idx:",idx)
        return
    }

    for c:= 0; c < n; c++{
        id1 := idx -c + n
        id2 := 2*n - idx - c - 1
        fmt.Println("idx:",idx)
        fmt.Println("id1:",id1,"id2:",id2,"c:",c)
        fmt.Println("cols:",cols,"d1:",d1,"d2:",d2)
        if !cols[c] && !d1[id1] && !d2[id2] {
            b:= make([]byte, n)
            for i:= range b{
                b[i] = '.'
            }
            b[c] = 'Q'
            board[idx] = string(b)
            fmt.Println("b:",string(b))
            fmt.Println("board:",board)
            cols[c], d1[id1],d2[id2] = true, true, true
            dfs(idx+1,n,cols,d1,d2,board,res)
            cols[c],d1[id1],d2[id2] = false,false,false
        }
    }
}

