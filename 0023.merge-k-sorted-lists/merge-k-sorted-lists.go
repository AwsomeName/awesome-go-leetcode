package problem0023

import "fmt"

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(lists []*ListNode) *ListNode {
    if lists == nil {return nil}
    var List_sort *ListNode
    for i := range(lists){
        List_sort = merge_two_sortedlist(List_sort, lists[i])
    }
    return List_sort
}


func merge_two_sortedlist(list_1 *ListNode, list_2 *ListNode) *ListNode{
    if list_1 == nil {return list_2}
    if list_2 == nil {return list_1}
    var DstList,tmpNode,tmp *ListNode

    fmt.Println(list_1.Val, list_2.Val)
    for list_1 != nil && list_2 != nil {
        fmt.Println(list_1.Val, list_2.Val)
        if list_1.Val < list_2.Val {
            tmpNode = list_1
            list_1 = list_1.Next
        }else {
            tmpNode = list_2
            list_2 = list_2.Next
        }
        if tmp != nil {tmp.Next = tmpNode} else {DstList = tmpNode}
        tmp = tmpNode
        tmpNode = tmpNode.Next
    }
    if list_1 != nil { tmp.Next = list_1}
    if list_2 != nil { tmp.Next = list_2}
    fmt.Println("Dstlist", DstList)
    return DstList
}
