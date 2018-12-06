package problem0037

import "fmt"
import "time"

type cell struct{ // encapsulates a single cell on a a Sudoku board
    Value int; // cell value 1..9 or 0 if unset
    NumPossibilities int; // number of possible (unconstrained) values
    Constraints [10]bool // if constraints[v] is true then value cant be v
}

var CELLS [9][9]cell
var bt [][]int

func set(i,j,v int) bool {
    c := &CELLS[i][j]
    if (c.Value == v){
        return true
    }
    if (c.Constraints[v]){
        return false
    }
    for i:=1;i<=9;i++{
        if  i!=v {
            c.Constraints[i] = true
        }
    }
    //c.Constraints[v] = false // and else need to be true
    c.NumPossibilities = 1
    c.Value = v
    //fmt.Println("Set func: set i,j:",i,j,"to:",v)
    //printCells()
    //fmt.Println("-----------")
    for k:=0; k<9; k++{
        //fmt.Println("prop to k,j:",k,j,v)
        if (i!=k && !updateConstraints(k,j,v)){
            return false
        }
        //fmt.Println("prop to i,k:",i,k,v)
        if (j != k && !updateConstraints(i,k,v)){
            return false
        }
        ix := (i/3)*3 + k/3
        jx := (j/3) * 3 + k%3
        //fmt.Println("prop to ix,jx:",ix,jx,v)
        if (ix != i && jx != j && !updateConstraints(ix,jx,v)){
            return false
        }
    }
    return true
}

func updateConstraints(i,j,v int) bool{
    c := &CELLS[i][j]
    if (c.Constraints[v]) {
        return true
    }
    if (c.Value == v) {
        return false
    }
    c.Constraints[v] = true
    c.NumPossibilities--
    if c.NumPossibilities > 1 {
        return true
    }
    //fmt.Println("NumPosib:",c.NumPossibilities)
    for x:= range c.Constraints{
        if !c.Constraints[x] {
            //fmt.Println("prop to set i,j,x:",i,j,x)
            return set(i,j,x)
        }
    }
    return false
}
func findValuesForEmptyCells() bool{
    for i:=0; i<9; i++{
        for j:=0; j<9; j++{
            if CELLS[i][j].Value == 0 {
                bt = append(bt, []int{i,j})
            }
        }
    }
    fmt.Println("bt len:",len(bt))
    for i := range bt{
        for j:= i+1; j < len(bt); j++{
            if CELLS[bt[i][0]][bt[i][1]].Value > CELLS[bt[j][0]][bt[j][1]].Value{
                bt[i],bt[j] = bt[j],bt[i]
            }
        }
    }
    if len(bt) > 0 {
        return backtrack(0)
    }else {
        return true
    }
}

func backtrack(k int) bool{
    if k >= len(bt) { return true}
    i := bt[k][0]
    j := bt[k][1]
    if CELLS[i][j].Value > 0 { return backtrack(k+1)}
    tmpConstraints := CELLS[i][j].Constraints
    var snapshot [9][9]cell
    snapshot = CELLS
    for v:=0; v <=9; v++{
        if !tmpConstraints[v]{
            if set(i,j,v) {
                if backtrack(k+1){
                    return true
                }
            }
            CELLS = snapshot
        }
    }
    return false
}

func printCells(){
    for i:= range CELLS{
        for j:= range CELLS[i]{
            fmt.Print(" ",CELLS[i][j].Value)
        }
        fmt.Println()
    }
}

func solveSudoku(board [][]byte) {
    start := time.Now()
    fmt.Println("start:",start)
    //fmt.Println("----------")
    for i:= range CELLS{
        for j:= range CELLS[i]{
            CELLS[i][j].NumPossibilities = 9
            CELLS[i][j].Value = 0
            CELLS[i][j].Constraints[0]= true
        }
    }
    fmt.Println("after init:",time.Now().Sub(start))

    for i:=0; i<9; i++{
        for j:=0; j<9; j++{
            //fmt.Println("set board[i][i]:",i,j,"to:",board[i][j])
            if board[i][j] != '.' && !set(i,j,int(board[i][j] - '0')){
                return
            }
            //fmt.Println("set done")
        }
    }
    //fmt.Println("===========init cells=========")
    //printCells()
    fmt.Println("after set:",time.Now().Sub(start))
    if !findValuesForEmptyCells(){
        return
    }
    fmt.Println("after cnt:",time.Now().Sub(start))

    for i:=0; i<9; i++{
        for j:=0; j<9; j++{
            if CELLS[i][j].Value > 0 {
                board[i][j] = byte(CELLS[i][j].Value + '0')
            }
        }
    }
    fmt.Println("after copy:",time.Now().Sub(start))
}
