package problem0102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	hdr, tail := len(stack), 0
	for hdr > tail {
		iitmp := make([]int, 0)
		idx := hdr
		for tail < idx {
			iitmp = append(iitmp, stack[tail].Val)
			if stack[tail].Left != nil {
				stack = append(stack, stack[tail].Left)
				hdr++
			}
			if stack[tail].Right != nil {
				stack = append(stack, stack[tail].Right)
				hdr++
			}
			tail++
		}
		ans = append(ans, iitmp)
	}
	return ans
}
