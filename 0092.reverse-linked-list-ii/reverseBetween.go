package problem0092

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	curCnt := 0

	var mBef, mIdx, pre *ListNode
	cur := head

	for cur != nil {
		fmt.Println(cur.Val)
		curCnt++
		if curCnt == m-1 {
			mBef = cur
			pre = cur
			cur = cur.Next
			continue
		}
		if curCnt == m {
			mIdx = cur
			pre = cur
			cur = cur.Next
			continue
		}
		if curCnt > m && curCnt < n {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
			continue
		}
		if curCnt == n {
			mIdx.Next = cur.Next
			if mBef != nil {
				mBef.Next = cur
			} else {
				head = cur
			}
			cur.Next = pre
			return head
		}
		cur = cur.Next
	}
	return head
}
