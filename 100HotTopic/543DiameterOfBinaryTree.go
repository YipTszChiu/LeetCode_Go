package _00HotTopic

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/?favorite=2cktkvj

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 树根不为空，直径至少是1
	res := 1

	// 路径长为路径经过的节点数减一，因此使用res作为全局变量记录最长路径
	var lenOfTree func(*TreeNode) int
	lenOfTree = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftLen := lenOfTree(root.Left)
		rightLen := lenOfTree(root.Right)
		// 此时的res仍然记录的是经过的节点数
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
