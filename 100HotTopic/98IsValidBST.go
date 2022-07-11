package _00HotTopic

import "math"

// 98. 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归思路
func isValidBST(root *TreeNode) bool {
	return judgeBST(root, math.MinInt64, math.MaxInt64)
}

func judgeBST(root *TreeNode, lower, upper int) bool {
	// 越过叶子节点返回true
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}
	// BST的子树仍然是BST，遍历左右子树；左子树的值不能大于当前根，右子树的值不能小于当前根
	return judgeBST(root.Left, lower, root.Val) && judgeBST(root.Right, root.Val, upper)
}

// 中序遍历思路
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	// 中序遍历
	for len(stack) > 0 || root != nil {
		// 往左子树走到底
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// BST中序遍历为递增序列，如果 root.Val 非递增，则不是BST
		if root.Val <= inorder {
			return false
		}
		// 更新保存中序遍历的上一个值
		inorder = root.Val
		// 走右子树
		root = root.Right
	}
	// 如果完整走完整棵树，说明符合BST要求
	return true
}
