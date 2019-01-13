package problem0110

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, x := check(root)
	return x
}

func check(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	if root.Left == nil && root.Right == nil {
		return 1, true
	}

	dl, tl := check(root.Left)
	dr, tr := check(root.Right)

	if !tl || !tr {
		return max(dl, dr) + 1, false
	} else {
		if dl-dr > 1 || dl-dr < -1 {
			return max(dl, dr) + 1, false
		}
		return max(dl, dr) + 1, true
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
