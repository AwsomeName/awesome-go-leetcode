package problem1022

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    if root == nil {
        return 0
    }

    cnt := make([]int,0)
    cntFromParent(root, 0, &cnt)

    sum := 0
    for _,x := range cnt {
        sum += x
        sum = sum % (1000000007)
    }
    // fmt.Println("cnt:", len(cnt))
    return sum
}

func cntFromParent(root *TreeNode, pre int, cnt *[]int) {
    // if root == nil {
    //     return
    // }
    // ans := 0
    if pre >= 1000000007 {
        pre = pre % 1000000007
    }
    cur := pre * 2 + root.Val
    // fmt.Println(root.Val)
    if root.Left == nil && root.Right == nil {
        // fmt.Println( cur)
        // cur = cur % (1000000007 * 2)
        *cnt = append(*cnt, cur)
        return
    }

    if root.Left != nil {
        cntFromParent(root.Left,  cur, cnt)
    }
    if root.Right != nil {
        cntFromParent(root.Right,  cur, cnt)
    }

    return
}
