package problem0024

type ListNode struct{
   Val int
   Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil { return nil }
    var new_head *ListNode
    if head.Next == nil { return head } else {
        new_head = head.Next
    }
    var pre_head *ListNode
    for head != nil && head.Next != nil {
        if pre_head != nil {
           pre_head.Next = head.Next
        }
        pre_head = head
        tmp := head.Next
        head.Next = head.Next.Next
        tmp.Next = head
        head = head.Next
    }
	return new_head
}
