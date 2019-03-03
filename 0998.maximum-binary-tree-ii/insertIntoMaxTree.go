package problem0998

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    // var newNode TreeNode
    newNode := TreeNode{val, nil, nil}
    if root == nil {
        return &newNode
    }

    if root.Val < val {
        newNode.Left = root
        return &newNode
    }else if root.Right == nil {
        root.Right = &newNode
        return root
    }

    preNode := root
    tmpNode := root.Right
    for tmpNode != nil {
        if val > tmpNode.Val {
            preNode.Right = &newNode
            newNode.Left = tmpNode
            return root
        }else{
            if tmpNode.Right == nil {
                tmpNode.Right = &newNode
                return root
            }else{
                preNode = tmpNode
                tmpNode = tmpNode.Right
            }
        }
    }

    return root
}
