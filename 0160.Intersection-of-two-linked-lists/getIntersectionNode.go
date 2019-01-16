package problem0160

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m, n := 0, 0
	p1, p2 := headA, headB

	for p1 != nil {
		m++
		p1 = p1.Next
	}

	for p2 != nil {
		n++
		p2 = p2.Next
	}

	p1, p2 = headA, headB

	if m > n {
		for i := 0; i < m-n; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i < n-m; i++ {
			p2 = p2.Next
		}
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1

}
