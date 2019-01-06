package problem0971

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	if root == nil || len(voyage) == 0 || root.Val != voyage[0] {
		return []int{-1}
	}
	vo := voyage
	cnt := 0
	ans := make([]int, 0)

	var chk func(*TreeNode) bool

	chk = func(root *TreeNode) bool {
		if root == nil || root.Val != vo[cnt] {
			return false
		}
		if root.Left != nil {
			if root.Left.Val == vo[cnt+1] {
				cnt++
				if !chk(root.Left) {
					return false
				}
			} else {
				root.Left, root.Right = root.Right, root.Left
				ans = append(ans, root.Val)
				if root.Left != nil {
					if root.Left.Val == vo[cnt+1] {
						cnt++
						if !chk(root.Left) {
							return false
						}
					} else {
						return false
					}
				}
			}
		}
		if root.Right != nil {
			if root.Right.Val == vo[cnt+1] {
				cnt++
				if !chk(root.Right) {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}

	if chk(root) {
		return ans
	} else {
		return []int{-1}
	}
}
