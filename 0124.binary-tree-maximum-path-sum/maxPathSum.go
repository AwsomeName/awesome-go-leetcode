package problem0124

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := root.Val
	cal(root, &ans)
	return ans
}

func cal(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left, right := max(0, cal(root.Left, ans)), max(0, cal(root.Right, ans))
	*ans = max(*ans, left+right+root.Val)
	return max(left, right) + root.Val
}

//     public int maxPathSum(TreeNode root) {
//         maxValue = Integer.MIN_VALUE;
//         maxPathDown(root);
//         return maxValue;
//     }

//     private int maxPathDown(TreeNode node) {
//         if (node == null) return 0;
//         int left = Math.max(0, maxPathDown(node.left));
//         int right = Math.max(0, maxPathDown(node.right));
//         maxValue = Math.max(maxValue, left + right + node.val);
//         return Math.max(left, right) + node.val;
//     }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
