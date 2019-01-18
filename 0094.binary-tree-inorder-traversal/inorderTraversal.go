package problem0094

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stk := make([]*TreeNode, 0)
	hdr := -1
	ans := make([]int, 0)

	for root != nil || hdr >= 0 {
		for root != nil {
			if hdr+1 <= len(stk)-1 {
				hdr++
				stk[hdr] = root
			} else {
				stk = append(stk, root)
				hdr++
			}
			root = root.Left
		}
		root = stk[hdr]
		hdr--
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}
