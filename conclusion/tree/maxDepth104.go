package tree

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max104(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

func max104(i, j int) int {
	if i > j {
		return i
	}
	return j
}
