package _00HotTopic

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归法
func isSymmetric(root *TreeNode) bool {
	return judgeSymmetric(root, root)
}

// 传入两个节点，判断其左右子节点是否对称
func judgeSymmetric(root1, root2 *TreeNode) bool {
	// 叶子节点
	if root1 == nil && root2 == nil {
		return true
	}
	// 一个为叶子节点另一个不是
	if root1 == nil || root2 == nil {
		return false
	}
	// 都不是叶子节点，首先判断这两个节点的值是否相等，然后递归判断
	return root1.Val == root2.Val && judgeSymmetric(root1.Left, root2.Right) && judgeSymmetric(root1.Right, root2.Left)
}
