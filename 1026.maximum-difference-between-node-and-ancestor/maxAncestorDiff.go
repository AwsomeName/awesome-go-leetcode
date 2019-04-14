package problem1026

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    if root == nil {
        return 0
    }

    ans := make([]int, 0)

    check(root, root.Val, root.Val, &ans)

    return max(ans)
}

func max(a []int) int{
    maxi := 0
    for _,x := range a {
        if maxi < x {
            maxi = x
        }
    }
    return maxi
}

func check(root *TreeNode, maxP, minP int, ans *[]int){
    if root == nil {
        return
    }

    if root.Val > maxP {
        maxP = root.Val
    }

    if root.Val < minP {
        minP = root.Val
    }

    newabs := maxP - minP
    if newabs < 0 {
        newabs = -newabs
    }

    *ans = append(*ans, newabs)


    if root.Left != nil {
        check(root.Left, maxP, minP, ans)
    }
    if root.Right != nil {
        check(root.Right, maxP, minP, ans)
    }
}
