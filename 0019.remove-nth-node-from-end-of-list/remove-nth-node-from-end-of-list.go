package problem0019

// ListNode 是节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	ptr_first, ptr_second := head, head
	for ptr_first.Next != nil {
		n--
		if n < 0 {
			ptr_second = ptr_second.Next
		}
		ptr_first = ptr_first.Next
	}
	if ptr_second.Next == nil {
		return nil
	}
	if n == 1 {
		return head.Next
	}
	ptr_second.Next = ptr_second.Next.Next
	return head
}
