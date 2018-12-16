package problem0958

import "fmt"

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	hdr, tal := 0, 0
	stk := make([]*TreeNode, 100)
	last := 0
	stk[hdr] = root
	hdr++
	k := 1
	cnt := 0
	for hdr > tal {
		tmp := stk[tal]
		if cnt == pow(2, k)-1 {
			k++
		}
		cnt++
		tal++
		if tmp.Val != last+1 && cnt != pow(2, k-1) {
			fmt.Println("tmp.Val:", tmp.Val, "last:", last, "cnt:", cnt, "k:", k, "j:", pow(2, k-1))
			return false
		}
		last = tmp.Val

		if cnt < pow(2, k)-1 && tal < hdr && (tmp.Left == nil || tmp.Right == nil) && (stk[tal].Left != nil || stk[tal].Right != nil) {
			return false
		}

		if tmp.Right == nil && tmp.Left != nil && (tmp.Left.Left != nil || tmp.Left.Right != nil) {
			return false
		}

		if tmp.Left == nil && tmp.Right != nil {
			return false
		}
		if tmp.Left != nil {
			stk[hdr] = tmp.Left
			hdr++
		}
		if tmp.Right != nil {
			stk[hdr] = tmp.Right
			hdr++
		}
	}
	return true
}

func pow(x, n int) int {
	ret := 1
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
