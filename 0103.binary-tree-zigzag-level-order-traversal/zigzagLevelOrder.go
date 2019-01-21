package problem0103

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// fmt.Println("-----------------------------")
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	hdr, tail := len(stack), 0

	flag := true
	for hdr > tail {
		fmt.Println(hdr, tail, len(stack))
		iitmp := make([]int, 0)
		idx, idx2 := hdr, hdr
		cnt := 0
		for flag && tail < idx {
			iitmp = append(iitmp, stack[idx2-1].Val)
			if stack[idx2-1].Left != nil {
				stack = append(stack, stack[idx2-1].Left)
				hdr++
			}
			if stack[idx2-1].Right != nil {
				stack = append(stack, stack[idx2-1].Right)
				hdr++
			}
			tail++
			idx2--
		}
		for !flag && tail < idx {
			iitmp = append(iitmp, stack[idx-1].Val)

			if stack[idx-1].Right != nil {
				stack = append(stack, stack[idx-1].Right)
				hdr++
			}
			if stack[idx-1].Left != nil {
				stack = append(stack, stack[idx-1].Left)
				hdr++
			}
			// tail ++
			idx--
			cnt++
		}
		tail += cnt
		ans = append(ans, iitmp)
		flag = !flag
	}
	return ans

}
