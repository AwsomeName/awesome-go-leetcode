package problem0095

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return genTrees(1, n)
}

func genTrees(start, end int) []*TreeNode {
	list := make([]*TreeNode, 0)
	if start > end {
		var tmp *TreeNode
		list = append(list, tmp)
		return list
	}
	if start == end {
		tmp := &TreeNode{start, nil, nil}
		list = append(list, tmp)
		return list
	}

	left, right := make([]*TreeNode, 0), make([]*TreeNode, 0)

	for i := start; i <= end; i++ {
		left = genTrees(start, i-1)
		right = genTrees(i+1, end)
		for _, x := range left {
			for _, y := range right {
				root := &TreeNode{i, x, y}
				list = append(list, root)
			}
		}
	}

	return list
}
