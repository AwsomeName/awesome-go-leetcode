package problem0037

import "fmt"

type cell struct{ // encapsulates a single cell on a a Sudoku board
    Value int; // cell value 1..9 or 0 if unset
    NumPossibilities int; // number of possible (unconstrained) values
    Constraints [10]bool // if constraints[v] is true then value cant be v
}

func(Cil cell) CellInit (){
    Cil.Value = 0
    Cil.NumPossibilities = 9
//    constraints
}

var CELLS [9][9]cell
var bt [][]int

// set the value of the cell to [v]
// the function also propagates constraints to other cells and deduce new value where possible
func set(i,j,v int) bool {
    // updating state of the cell
    c := CELLS[i][j]
    if (c.Value == v){
        return true
    }
    if (c.Constraints[v]){
        return false
    }
    c.Constraints[v] = false // and else need to be true
    c.NumPossibilities = 1
    c.Value = v
    // propagating constraints
    for k:=0; k<9; k++{
        // to the row:
        if (i!=k && !updateConstraints(k,j,v)){
            return false
        }
        // to the column:
        if (j != k && !updateConstraints(i,k,v)){
            return false
        }
        // to the 3*3 square
        ix := (i/3)*3 + k/3
        jx := (j/3) * 3 + k%3
        if (ix != i && jx != j && !updateConstraints(ix,jx,v)){
            return false
        }
    }
    return true
}

// update constraints of the call i,j by excluding possibility of excludedValue
// once there's one possibility left the function recurses back into set()
func updateConstraints(i,j,v int) bool{
    c := CELLS[i][j]
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
    for _,cons := range c.Constraints{
        if !cons {
            return set(i,j,v)
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






    return true
}


func solveSudoku(board [][]byte) {
    fmt.Println("----------")
    // declear 9*9 cells and init
    for i:= range CELLS{
        for j:= range CELLS[i]{
            CELLS[i][j].CellInit()
            fmt.Println("cell.value:",CELLS[i][i].Value)
        }
    }

    for i:=0; i<9; i++{
        for j:=0; j<9; j++{
            if board[i][j] != '.' && !set(i,j,int(board[i][j] - '0')){
                return
            }
        }
    }
    if !findValuesForEmptyCells(){
        return
    }

    for i:=0; i<9; i++{
        for j:=0; j<9; j++{
            if CELLS[i][j].Value > 0 {
                board[i][j] = byte(CELLS[i][j].Value + '0')
            }
        }
    }
}
