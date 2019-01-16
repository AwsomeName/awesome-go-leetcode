package problem0108

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func sortedArrayToBST(nums []int) *TreeNode {
//     n := len(nums)
//     if n==0 {
//         return nil
//     }
//     mid := n/2
//     var head TreeNode
//     head.Val = nums[mid]
//     head.Left = sortedArrayToBST(nums[:mid])
//     head.Right = sortedArrayToBST(nums[mid+1:])

//     return &head
// }

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	} else if len(nums) == 2 {
		leftNode := TreeNode{Val: nums[0], Left: nil, Right: nil}
		return &TreeNode{Val: nums[1], Left: &leftNode, Right: nil}
	} else {
		p := 0 + (len(nums)-1)/2
		return &TreeNode{Val: nums[p], Left: sortedArrayToBST(nums[:p]), Right: sortedArrayToBST(nums[p+1 : len(nums)])}
	}
}
