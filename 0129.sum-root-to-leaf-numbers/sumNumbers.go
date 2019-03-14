package problem0129

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return root.Val
    }
    ans := 0
    if root.Left != nil {
        sum(root.Left, root.Val, &ans)
    }
    if root.Right != nil {
        sum(root.Right, root.Val, &ans)
    }
    return ans
}

func sum(root *TreeNode, psum int, ans *int){
    if root.Left == nil && root.Right == nil {
        *ans += psum*10 + root.Val
        return
    }
    if root.Left != nil {
        sum(root.Left, psum*10 + root.Val, ans)
    }
    if root.Right != nil {
        sum(root.Right, psum*10 + root.Val, ans)
    }
}

