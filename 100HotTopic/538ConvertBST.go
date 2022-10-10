package _00HotTopic

// 538. 把二叉搜索树转换为累加树
// https://leetcode.cn/problems/convert-bst-to-greater-tree/?favorite=2cktkvj

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 逆中序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func convertBST(root *TreeNode) *TreeNode {
	// 思路：二叉搜索树左右边的节点是最大的，而题干要求大于等于当前节点的和，即其右子树和+当前节点
	// 因此通过右-根-左的逆中序遍历计算和即可
	var dfs func(*TreeNode)
	sum := 0
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			sum += root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}
