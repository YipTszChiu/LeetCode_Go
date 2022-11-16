package tree

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/description/?favorite=2cktkvj

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0

	// 在正常的计算树高函数中插入对直径的计算
	var lenOfTree func(*TreeNode) int
	lenOfTree = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftLen := lenOfTree(root.Left)
		rightLen := lenOfTree(root.Right)
		res = max543(res, leftLen+rightLen+1)
		return max543(leftLen, rightLen) + 1
	}

	lenOfTree(root)

	return res - 1
}

func max543(i, j int) int {
	if i > j {
		return i
	}
	return j
}
