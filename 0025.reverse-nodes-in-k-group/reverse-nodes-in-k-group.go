package problem0025

import "github.com/aQuaYi/LeetCode-in-Go/kit"

type ListNode = kit.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}
	dst_head := head
	for i := 0; i < k-1; i++ {
		if dst_head.Next != nil {
			dst_head = dst_head.Next
		} else {
			return head
		}
	}
	rev_cnt := 0
	var prehead, tmphead *ListNode
	for head != nil {
		if rev_cnt == 0 {
			tmphead = head
			head = head.Next
			rev_cnt++
			if head == nil && prehead != nil {
				prehead.Next = tmphead
			} else if head == nil && prehead == nil {
				return tmphead
			}
		} else if rev_cnt == k-1 || head.Next == nil {
			if rev_cnt < k-1 {
				if prehead != nil {
					prehead.Next = tmphead
					break
				} else {
					return tmphead
				}
			}
			rev_cnt = 0
			///这里，截取一段做反转，并保存，将新头连到上一个尾巴。
			tmp := head.Next
			head.Next = nil
			tmph, tmpt := reverseListNode(tmphead)
			if prehead != nil {
				prehead.Next = tmph
			}
			prehead = tmpt
			head = tmp
		} else {
			head = head.Next
			rev_cnt++
		}
	}
	return dst_head
}

func reverseListNode(head *ListNode) (*ListNode, *ListNode) {
	var pre_head, new_end *ListNode
	new_end = head
	for head != nil {
		tmp := head.Next
		head.Next = pre_head
		pre_head = head
		head = tmp
	}
	return pre_head, new_end
}
