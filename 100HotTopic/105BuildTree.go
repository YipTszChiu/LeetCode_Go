package _00HotTopic

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归方法
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 先序遍历的首位作为根节点
	root := &TreeNode{preorder[0], nil, nil}
	// 在中序遍历中找到该根节点
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	// 按照左右子树划分
	root.Left = buildTree(preorder[1:len(inorder[:i+1])], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i+1]):], inorder[i+1:])

	return root
}
