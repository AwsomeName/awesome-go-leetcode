package problem0114

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	adjust(root)
	reverse(root)
}

func adjust(root *TreeNode) {
	if root == nil {
		return
	}
	var pre *TreeNode
	stack := make([]*TreeNode, 100)
	hdr := 0
	cur := root
	for cur != nil || hdr > 0 {
		if cur == nil {
			if hdr > 0 {
				cur = stack[hdr-1]
				hdr--
			} else {
				return
			}
		}
		if pre != nil {
			pre.Left = cur
			pre.Right = nil
		}
		pre = cur
		if cur.Right != nil {
			stack[hdr] = cur.Right
			hdr++
		}
		cur = cur.Left
	}
}

func reverse(root *TreeNode) {
	for root != nil {
		root.Right = root.Left
		root.Left = nil
		root = root.Right
	}
}
