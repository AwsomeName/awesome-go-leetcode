package problem0023

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return merge_two_sortedlist(lists[0], lists[1])
	} else if len(lists) > 2 {
		half := len(lists) / 2
		return merge_two_sortedlist(merge(lists[:half]), merge(lists[half:]))
	}
	return nil
}

func merge_two_sortedlist(list_1 *ListNode, list_2 *ListNode) *ListNode {
	if list_1 == nil {
		return list_2
	}
	if list_2 == nil {
		return list_1
	}
	var DstList, tmpNode, tmp *ListNode

	for list_1 != nil && list_2 != nil {
		if list_1.Val < list_2.Val {
			tmpNode = list_1
			list_1 = list_1.Next
		} else {
			tmpNode = list_2
			list_2 = list_2.Next
		}
		if tmp != nil {
			tmp.Next = tmpNode
		} else {
			DstList = tmpNode
		}
		tmp = tmpNode
		tmpNode = tmpNode.Next
	}
	if list_1 != nil {
		tmp.Next = list_1
	}
	if list_2 != nil {
		tmp.Next = list_2
	}
	return DstList
}
