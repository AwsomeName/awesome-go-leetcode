package problem1008

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    var root TreeNode
    if len(preorder) == 0 {
        return nil
    }
    root.Val = preorder[0]
    index := 0
    for i,x := range preorder{
        if x > root.Val {
            index = i
            break
        }
    }
    fmt.Println(index, preorder)
    if index > 0 {
        root.Left = bstFromPreorder(preorder[1:index])
    }else if index == 0 && len(preorder)>1 {
        root.Left = bstFromPreorder(preorder[1:])
    }

    if index > 0 && index < len(preorder){
        root.Right = bstFromPreorder(preorder[index:])
    }


    return &root
}
