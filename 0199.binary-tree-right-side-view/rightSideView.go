package problem0199

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil { return []int{}}

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

    ans := make([]int,0)
    for i:= 0; i < len(stk); i++{
        if len(stk[i]) > 0 && stk[i][len(stk[i])-1] != nil{
            ans = append(ans, stk[i][len(stk[i])-1].Val)
        }
    }
    return ans
}
