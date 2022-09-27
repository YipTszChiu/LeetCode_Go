package _00HotTopic

// 337. 打家劫舍 III
// https://leetcode.cn/problems/house-robber-iii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 树形动态规划
// 思路：每一个节点可以分为 selected 和 notSelected 两种状态
// 1.如果节点被选中，那么他的左右子节点不能被选中，所以最大权值为：
//  selected(node) = notSelected(node.left) + notSelected(node.right)
// 2.如果节点不被选中，那么他的左右子节点可以选中也可以不选中，所以最大权值为：
//  notSelected(node) = max(selected(node.left) + notSelected(node.left)) + max(selected(node.right) + notSelected(node.right)

// 该dfs相当于后序遍历，返回值 []int{a, b}：a表示该节点选中的最大权值，b表示该节点不被选中的最大权值
func rob337Dfs(node *TreeNode) []int {
	// 空节点不提供权值
	if node == nil {
		return []int{0, 0}
	}
	// 递归调用dfs返回左右子节点选中or不选中的最大权值
	l, r := rob337Dfs(node.Left), rob337Dfs(node.Right)
	// 当前节点被选中，则左右子节点不能被选中
	selected := node.Val + l[1] + r[1]
	// 当前节点不选中
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func rob337(root *TreeNode) int {
	money := rob337Dfs(root)
	return max(money[0], money[1])
}
