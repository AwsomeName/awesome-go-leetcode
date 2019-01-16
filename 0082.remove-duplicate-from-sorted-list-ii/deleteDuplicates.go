package problem0082

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newhdr *ListNode
	var pre *ListNode

	flag := false
	for head != nil {
		if head.Next != nil && head.Next.Val == head.Val {
			flag = true
			for head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
		}
		if flag {
			flag = false
		} else {
			if pre == nil {
				// pre = newhdr
				// pre.Val = head.Val
				pre = &ListNode{head.Val, nil}
				// pre.Next = p
				newhdr = pre
				// pre = p
			} else {
				p := &ListNode{head.Val, nil}
				pre.Next = p
				pre = p
			}
		}
		head = head.Next
	}
	return newhdr

}
