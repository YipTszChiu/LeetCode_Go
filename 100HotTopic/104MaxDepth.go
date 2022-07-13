package _00HotTopic

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 层次遍历
func maxDepth(root *TreeNode) int {
	// 空值处理
	if root == nil {
		return 0
	}
	// 初始化高度
	level := 0
	// 初始化层次遍历使用的队列，并且将root入队
	queue := []*TreeNode{root}
	// last指向每层最后一个节点
	last := root
	// 进行层次遍历
	for len(queue) > 0 {
		// 出队
		node := queue[0]
		// node 左右子节点入队
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		// 判断是否为每层最后一个节点
		if node == last {
			// 移动 last
			last = queue[len(queue)-1]
			// 层数加1
			level++
		}
		// node 出队
		queue = queue[1:]
	}

	return level

}

// 深度优先遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)+1)
}
