package problem0025

import "github.com/aQuaYi/LeetCode-in-Go/kit"
import "fmt"

type ListNode = kit.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k==1 || head == nil {
        return head
    }
    dst_head := head
    for i:=0; i<k-1; i++{
        dst_head = dst_head.Next
    }
    rev_cnt := 0
    var prehead, tmphead *ListNode
    for head != nil {
        fmt.Println( head.Val )
        if rev_cnt == 0{
            tmphead = head
            head = head.Next
            rev_cnt++
        }else if rev_cnt == k-1 || head.Next == nil {
            if rev_cnt < k-1 {
               prehead.Next = tmphead
               break
            }
            rev_cnt = 0
            ///这里，截取一段做反转，并保存，将新头连到上一个尾巴。
            tmp:= head.Next
            head.Next = nil
            fmt.Println("tmphead,VAl: ", tmphead.Val)
            tmph,tmpt:= reverseListNode(tmphead )
            fmt.Println("tmph.val: ",tmph.Val)
            fmt.Println("tmpt.val: ",tmpt.Val)
            if prehead != nil { prehead.Next = tmph}
            prehead = tmpt
            head = tmp
        }else {
            head = head.Next
            rev_cnt++
        }
        fmt.Println("rev_cnt: ", rev_cnt)
    }
    t := dst_head
    fmt.Println(" print result: ")
    for t != nil {
        fmt.Println(t.Val)
        t = t.Next
    }
    return dst_head
}

func reverseListNode(head *ListNode) (*ListNode, *ListNode){
    var pre_head,new_end *ListNode
    new_end = head
    for head != nil{
        tmp:= head.Next
        head.Next = pre_head
        pre_head = head
        head = tmp
    }
    return pre_head,new_end
}
