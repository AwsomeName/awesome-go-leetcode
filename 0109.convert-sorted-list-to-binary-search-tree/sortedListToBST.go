package problem0109

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	var root *TreeNode
	if head == nil {
		return root
	}
	root = &TreeNode{0, nil, nil}

	var fast, slow, sh *ListNode
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fmt.Println(fast.Val)
		fast = fast.Next.Next
		sh = slow
		slow = slow.Next
	}

	root.Val = slow.Val
	if sh != nil {
		sh.Next = nil
	}
	if slow != head {
		root.Left = sortedListToBST(head)
	}
	if slow.Next != nil {
		root.Right = sortedListToBST(slow.Next)
	}

	return root
}
