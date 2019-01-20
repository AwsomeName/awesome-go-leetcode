package problem0979

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type nT struct {
	Need  int
	Have  int
	Left  *nT
	Right *nT
}

func distributeCoins(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cnt := 0

	nRoot := &nT{0, 0, nil, nil}
	build(nRoot, root)

	_, _ = setNT(nRoot, root)

	cnt = count(nRoot)

	return cnt
}

func setNT(nRoot *nT, root *TreeNode) (int, int) {
	if nRoot == nil {
		return 0, 0
	}
	needLeft, haveLeft := setNT(nRoot.Left, root.Left)
	needRight, haveRight := setNT(nRoot.Right, root.Right)

	nRoot.Need = 1 + needLeft + needRight
	nRoot.Have = root.Val + haveLeft + haveRight
	return nRoot.Need, nRoot.Have
}

func count(nRoot *nT) int {
	if nRoot == nil {
		return 0
	}
	cnt := abs(nRoot.Have, nRoot.Need)
	cnt += count(nRoot.Left)
	cnt += count(nRoot.Right)
	return cnt
}

func build(nRoot *nT, root *TreeNode) {
	if root.Left != nil {
		left := &nT{0, 0, nil, nil}
		nRoot.Left = left
		build(nRoot.Left, root.Left)
	}

	if root.Right != nil {
		Right := &nT{0, 0, nil, nil}
		nRoot.Right = Right
		build(nRoot.Right, root.Right)
	}
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
