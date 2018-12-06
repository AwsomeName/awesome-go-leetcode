package problem0951

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flipEq(root1 *TreeNode, root2 *TreeNode) bool{
    if root1 == nil && root2 == nil { return true}
    if root1 == nil || root2 == nil { return false}
    if root1.Val != root2.Val { return false }
//    if root1.Left == nil && root2.Left == nil && root1.Right == nil && root2.Right == nil { return true}
    return flipEq(root1.Left,root2.Left)&&flipEq(root1.Right,root2.Right)||flipEq(root1.Left,root2.Right)&&flipEq(root1.Right,root2.Left)
}
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool{
    return flipEq(root1, root2)
}
