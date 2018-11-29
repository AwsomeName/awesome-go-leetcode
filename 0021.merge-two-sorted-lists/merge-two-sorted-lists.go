package problem0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var ptr_head, ptr_last *ListNode
	if l1.Val < l2.Val {
		ptr_head = l1
		l1 = l1.Next
	} else {
		ptr_head = l2
		l2 = l2.Next
	}
	ptr_last = ptr_head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr_tmp := l1
			ptr_last.Next = ptr_tmp
			l1 = l1.Next
		} else {
			ptr_tmp := l2
			ptr_last.Next = ptr_tmp
			l2 = l2.Next
		}
		ptr_last = ptr_last.Next
	}
	if l1 != nil {
		ptr_last.Next = l1
	} else {
		ptr_last.Next = l2
	}

	return ptr_head
}
