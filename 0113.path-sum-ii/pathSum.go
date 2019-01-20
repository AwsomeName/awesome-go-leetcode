package problem0113

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func pathSum(root *TreeNode, sum int) [][]int {
//     ans := make([][]int,0)
//     if root == nil {return ans}
//     fmt.Println("---------------------------------------")
//     check(root, &ans, []int{}, 0, sum)
//     return ans
// }

// func check(root *TreeNode, ans *[][]int, tmplist []int, tmpSum, sum int){
//     fmt.Println(root.Val, tmplist, tmpSum)
//     if root == nil {
//         return
//     }
//     tmp := make([]int,len(tmplist))
//     copy(tmp, tmplist)

//     if root.Val + tmpSum == sum {
//         if root.Left == nil && root.Right == nil {
//             tmp = append(tmp, root.Val)
//             *ans = append(*ans, tmp)
//             return
//         }
//     }
//     if root.Left != nil {
//         tmp = append(tmp, root.Val)
//         check(root.Left, ans, tmp, tmpSum + root.Val, sum)
//     }
//     if root.Right != nil {
//         tmplist = append(tmplist, root.Val)
//         check(root.Right, ans, tmplist, tmpSum+root.Val, sum)
//     }
//     return
// }

func pathSum(root *TreeNode, sum int) [][]int {

	ret := make([][]int, 0)
	path := make([]int, 0)
	hasPathSum(root, sum, path, &ret)

	return ret
}

func hasPathSum(root *TreeNode, sum int, path []int, ret *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		path = append(path, root.Val)
		*ret = append(*ret, append([]int(nil), path...))
		return
	}

	hasPathSum(root.Left, sum-root.Val, append(path, root.Val), ret)
	hasPathSum(root.Right, sum-root.Val, append(path, root.Val), ret)
}
