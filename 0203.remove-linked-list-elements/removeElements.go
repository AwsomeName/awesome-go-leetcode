package problem0203

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head != nil && head.Next != nil {
		head.Next = removeElements(head.Next, val)
	}
	return head
}
