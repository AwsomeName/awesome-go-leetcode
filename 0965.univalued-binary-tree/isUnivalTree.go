package problem0965

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	val := root.Val
	var chk func(cur *TreeNode, val int) bool
	chk = func(cur *TreeNode, val int) bool {
		if cur.Val != val {
			return false
		}
		if cur.Left != nil {
			if !chk(cur.Left, val) {
				return false
			}
		}
		if cur.Right != nil {
			if !chk(cur.Right, val) {
				return false
			}
		}
		return true
	}

	return chk(root, val)
}
