package _00HotTopic

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func inorderTraversal1(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		// 越过叶子节点返回
		if root == nil {
			return
		}
		// 左
		inorder(root.Left)
		// 根
		ans = append(ans, root.Val)
		// 右
		inorder(root.Right)
	}

	inorder(root)

	return ans
}

// 非递归法
func inorderTraversal2(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	// 初始化栈
	stack := []*TreeNode{}
	// 中序遍历
	for root != nil || len(stack) > 0 {
		// 左子节点入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 当前节点遍历到空节点，说明越过叶子节点，此时出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		// 遍历右子节点
		root = root.Right
	}

	return
}
