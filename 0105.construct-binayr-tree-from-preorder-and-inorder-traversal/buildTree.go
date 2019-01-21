package problem0105

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := len(preorder)
	var root *TreeNode
	if m == 0 {
		return root
	}
	root = &TreeNode{preorder[0], nil, nil}
	idx := 0
	for i, x := range inorder {
		if x == root.Val {
			idx = i
			break
		}
	}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
