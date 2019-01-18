package problem0086

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var left, right, lHdr, rHdr *ListNode
	if head == nil {
		return nil
	}

	for head != nil {
		if head.Val < x {
			if lHdr == nil {
				left = &ListNode{head.Val, nil}
				lHdr = left
			} else {
				left.Next = &ListNode{head.Val, nil}
				left = left.Next
			}
		} else {
			if rHdr == nil {
				right = &ListNode{head.Val, nil}
				rHdr = right
			} else {
				right.Next = &ListNode{head.Val, nil}
				right = right.Next
			}
		}
		head = head.Next
	}
	if lHdr == nil {
		return rHdr
	}

	left.Next = rHdr

	return lHdr
}
