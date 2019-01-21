package problem0234

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	hdr := -1
	stack := make([]int, 0)

	cur := head
	for cur != nil { //&& (hdr == -1 || cur.Val != stack[hdr]) {
		stack = append(stack, cur.Val)
		hdr++
		cur = cur.Next
	}
	flag := true
	for i := 0; i < len(stack)/2; i++ {
		if stack[i] == stack[len(stack)-i-1] {
			flag = true
			// continue
		} else {
			// fmt.Println(stack[i], stack[len(stack)-i-1],stack, i)
			return false
		}
	}

	return flag
}
