package _00HotTopic

import "math"

// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归计算贡献值
// 时间复杂度O(N): N是二叉树节点个数，对每个节点访问不会超过两次
// 空间复杂度O(N): N是二叉树节点个数，递归调用层数等于二叉树高度，最坏情况下高度 = 二叉树节点个数
func maxPathSum(root *TreeNode) int {
	// 思路：叶子节点贡献值为叶子节点本身的 val
	// 节点最大路径和 = 左、右子节点的最大贡献值 + 该节点 val
	// 使用递归回溯到root，用变量 maxSum 记录最大路径和
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 递归调用 maxGain
		// 只有贡献值大于0才会计入最大路径和
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 节点最大路径和 = 左、右子节点的最大贡献值 + 该节点 val
		nodePathSum := node.Val + leftGain + rightGain
		// 更新最大路径和
		maxSum = max(nodePathSum, maxSum)
		// 向上返回当前节点的最大贡献值
		// 注意此处和 nodePathSum 最大路径和的区别：
		// 由于以当前节点来看，可以走左→根→右的路径
		// 但向上递归时，父节点走到当前节点，再向下走时只能选择左 or 右其中一条路径
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
