package problem0099

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func recoverTree(root *TreeNode)  {
//     if root == nil {return }
//     if root.Left == nil && root.Right == nil { return }

//     nums := make([]int,0)

//     search(root, &nums)

//     right := make([]int, len(nums))
//     copy(right, nums)
//     sort.Ints(right)

//     idx := 0
//     for i:= 0; i < len(nums); i++{
//         if nums[i] !=  right[i] {
//             idx = i
//             break
//         }
//     }
//     fmt.Println(nums, idx)
//     set(root, nums[idx], right[idx])

//     return
// }

// func set(root *TreeNode, a, b int){
//     if root == nil { return }
//     if root.Val == a {
//         root.Val = b
//     }else if root.Val == b {
//         root.Val = a
//     }
//     set(root.Left, a, b)
//     set(root.Right, a, b)
// }

// func search(root *TreeNode, nums *[]int){
//     if root == nil {return}
//     search(root.Left, nums)
//     *nums = append(*nums, root.Val)
//     search(root.Right, nums)
// }

func checkLeft(root *TreeNode) *TreeNode {
	cur := root
	var res, last *TreeNode
	for cur != nil {
		if cur.Left == nil {
			if last != nil && cur.Val <= last.Val {
				res = cur
			}
			cur, last = cur.Right, cur
		} else {
			tmp := cur.Left
			for tmp.Right != nil && tmp.Right != cur {
				tmp = tmp.Right
			}
			if tmp.Right == cur {
				tmp.Right = nil
				if last != nil && cur.Val <= last.Val {
					res = cur
				}
				cur, last = cur.Right, cur
			} else {
				tmp.Right = cur
				cur = cur.Left
			}
		}
	}
	return res
}
func checkRight(root *TreeNode) *TreeNode {
	cur := root
	var res, last *TreeNode
	for cur != nil {
		if cur.Right == nil {
			if last != nil && cur.Val >= last.Val {
				res = cur
			}
			cur, last = cur.Left, cur
		} else {
			tmp := cur.Right
			for tmp.Left != nil && tmp.Left != cur {
				tmp = tmp.Left
			}
			if tmp.Left == cur {
				tmp.Left = nil
				if last != nil && cur.Val >= last.Val {
					res = cur
				}
				cur, last = cur.Left, cur
			} else {
				tmp.Left = cur
				cur = cur.Right
			}
		}
	}
	return res
}

func recoverTree(root *TreeNode) {
	t1, t2 := checkLeft(root), checkRight(root)
	t1.Val, t2.Val = t2.Val, t1.Val
}
