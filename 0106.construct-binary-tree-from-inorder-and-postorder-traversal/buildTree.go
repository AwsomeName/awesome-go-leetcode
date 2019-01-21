package problem0106

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	m := len(postorder)
	var root *TreeNode
	if m == 0 {
		return root
	}
	root = &TreeNode{postorder[m-1], nil, nil}
	idx := 0
	for i, x := range inorder {
		if x == root.Val {
			idx = i
			break
		}
	}
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:m-1])
	return root
}
