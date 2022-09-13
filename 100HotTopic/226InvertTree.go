package _00HotTopic

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/?favorite=2cktkvj

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	// 空值处理
	if root == nil {
		return root
	}
	// 交换左右子树
	root.Left, root.Right = root.Right, root.Left
	// 对下层进行交换
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
