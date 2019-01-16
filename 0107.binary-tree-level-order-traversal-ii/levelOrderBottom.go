package problem0107

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil { return [][]int{}}

    stk := make([][]*TreeNode,0)

    rootmp := []*TreeNode{root}
    stk = append(stk, rootmp)

    for len(stk[len(stk)-1]) > 0 {
        tmp := make([]*TreeNode,0)
        for i:= range stk[len(stk)-1]{
            if stk[len(stk)-1][i].Left != nil{
                tmp = append(tmp, stk[len(stk)-1][i].Left)
            }
            if stk[len(stk)-1][i].Right != nil{
                tmp = append(tmp, stk[len(stk)-1][i].Right)
            }
        }
        stk = append(stk, tmp)
    }

    ans := make([][]int,0)
    for i:= len(stk)-1; i >=0; i--{
        tmp := make([]int,0)//,len(stk[i]))
        for j:= range stk[i]{
            if stk[i][j] != nil{
                tmp = append(tmp,stk[i][j].Val)
            }
        }
        if len(tmp) > 0{
            ans = append(ans, tmp)
        }
    }
    return ans
}
