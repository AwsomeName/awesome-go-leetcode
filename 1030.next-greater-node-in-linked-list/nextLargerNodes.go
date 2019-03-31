package problem1030

func nextLargerNodes(head *ListNode) []int {
    ans := make([]int,0)
    if head == nil {
        return []int{}
    }

    for head != nil {
        ans = append(ans, findNext(head))
        head = head.Next
    }

    return ans
}

func findNext(head *ListNode) int{
    if head == nil {
        return 0
    }
    // ans := 0
    tmp := head.Val
    head = head.Next
    for head != nil{

        if tmp < head.Val{
            return head.Val
        }else{
            head = head.Next
        }
    }
    return 0
}
