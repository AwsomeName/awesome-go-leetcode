package problem0083

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := head
	for tmp != nil && tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
			continue
		}
		tmp = tmp.Next
	}
	return head
}
