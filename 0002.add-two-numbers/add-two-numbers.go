package problem0002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * Val int
 * Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*var ans *ListNode*/
	ans := &ListNode{
		Val: 0,
	}
	tmp_ans := ans
	var ok bool
	ok = true
	var c_bit int
	var ok_1, ok_2 int
	for ok {

		/*if l1 != nil && tmp_ans != nil{*/
		if l1 != nil {
			tmp_ans.Val += l1.Val
		}
		if l2 != nil {
			tmp_ans.Val += l2.Val
		}
		if c_bit > 0 {
			tmp_ans.Val += 1
		}
		if tmp_ans.Val >= 10 {
			c_bit = 1
			tmp_ans.Val -= 10
		} else {
			c_bit = 0
		}

		if l1 != nil && l1.Next != nil {
			ok_1 = 1
			l1 = l1.Next
		} else {
			ok_1 = 0
			l1 = nil
		}
		if l2 != nil && l2.Next != nil {
			ok_2 = 1
			l2 = l2.Next
		} else {
			ok_2 = 0
			l2 = nil
		}
		if ok_1+ok_2 == 0 {
			ok = false
			if c_bit > 0 {
				tmp_ans.Next = &ListNode{Val: 1}
			}
		} else {
			tmp_ans.Next = &ListNode{Val: 0}
			tmp_ans = tmp_ans.Next
		}
	}
	return ans
}
