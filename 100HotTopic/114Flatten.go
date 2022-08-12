package _00HotTopic

// 114. 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一：存储先序遍历的结果，重构左右子树
// 时间复杂度：O(n) 先序遍历O(n) + 重构子树O(n) → O(n)
// 空间复杂度：O(n) 取决于二叉树节点数
func flatten(root *TreeNode) {
	// 进行先序遍历
	preOrderList := preOrder(root)
	// 根据先序遍历的结果将树重构
	for i := 0; i < len(preOrderList)-1; i++ {
		node, next := preOrderList[i], preOrderList[i+1]
		node.Left, node.Right = nil, next
	}
}

func preOrder(node *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if node != nil {
		list = append(list, node)
		list = append(list, preOrder(node.Left)...)
		list = append(list, preOrder(node.Right)...)
	}
	return list
}

// 解法二：寻找前驱节点
// 时间复杂度：O(n) 取决于节点数量
// 空间复杂度：O(1) 原地转换
func flatten2(root *TreeNode) {
	// 思路：当前节点的左子节点空，按照先序遍历不需要操作；如果不空，访问完左子树最后一个节点后（前驱节点），访问该节点的右子节点
	// 问题转换为寻找前驱节点
	curr := root
	// 当前节点不空则持续将树转为链
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			// pre记录前驱节点、持续向右子树寻找
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			// 确认前驱节点，将当前节点链接到前驱节点的右子节点
			pre.Right = curr.Right
			// 当前节点左子节点按题目要求置为空，右子节点按照先序遍历置为 next
			curr.Left, curr.Right = nil, next
		}
		// 当前节点左子节点为空，直接遍历右子节点
		curr = curr.Right
	}
}
