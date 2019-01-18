package problem0206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var cur, pre *ListNode
	pre = head
	cur = head.Next
	pre.Next = nil

	for cur != nil {
		// if pre == nil{
		//     pre = cur
		//     cur = cur.Next
		// }else{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		// }
	}
	return pre
}
