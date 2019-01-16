package problem0112

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	// if root.Val == sum && (root.Left != nil && root.Right != nil){
	//     return false
	// }
	tmp := sum - root.Val
	return hasPathSum(root.Left, tmp) || hasPathSum(root.Right, tmp)
}
