package _00HotTopic

// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) (ans [][]int) {
	// 空值处理
	if root == nil {
		return nil
	}
	// 初始化队列，root 入队
	queue := []*TreeNode{root}
	// tag 指向每层最后一个位置
	tag := root
	arr := []int{}
	// 进行层次遍历
	for len(queue) > 0 {
		// 队首出队
		node := queue[0]
		// 记录数值
		arr = append(arr, node.Val)
		// 队首节点的左右子节点入队
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		// 已经遍历到个该层最后一个节点，更新tag，返回 arr
		if node == tag {
			tag = queue[len(queue)-1]
			ans = append(ans, arr)
			arr = []int{}
		}
		// 出队
		queue = queue[1:]
	}

	return
}
