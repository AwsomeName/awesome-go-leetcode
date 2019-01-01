package problem0100

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	if p.Val != q.Val || p.Left == nil && q.Left != nil || p.Right == nil && q.Right != nil || q.Left == nil && p.Left != nil || q.Right == nil && q.Right != nil {
		return false
	}

	if p.Left != nil && !isSameTree(p.Left, q.Left) {
		return false
	}

	if p.Right != nil && !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
