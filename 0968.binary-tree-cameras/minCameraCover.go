package problem0968

func minCameraCover(root *TreeNode) int {
	ans := 0
	tmp := dfs(root, &ans)
	if tmp < 1 {
		return 1 + ans
	}
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	left, right := 0, 0
	if root.Left == nil {
		left = 2
	} else {
		left = dfs(root.Left, ans)
	}
	if root.Right == nil {
		right = 2
	} else {
		right = dfs(root.Right, ans)
	}

	if left == 0 || right == 0 {
		*ans++
		return 1
	}

	if left == 1 || right == 1 {
		return 2
	}

	return 0
}
