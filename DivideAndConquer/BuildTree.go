package DivideAndConquer

//剑指 Offer 07. 重建二叉树
//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	//空值处理
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//从先序遍历获得树根
	root := &TreeNode{
		preorder[0],
		nil,
		nil,
	}
	//中序遍历的索引，用于在中序序遍历中找到当前子树的树根
	inorderIndex := 0
	for ; inorderIndex < len(inorder); inorderIndex++ {
		//当在中序遍历中找到先序遍历的树根，跳出循环
		if inorder[inorderIndex] == preorder[0] {
			break
		}
	}

	//递归调用
	//从inorder[inorderIndex]这一位开始：
	//左边为左子树的内容，即inorder[:inorderIndex]；而对preorder来说，要计算中序遍历0~inorderIndex节点的个数
	root.Left = buildTree(preorder[1:len(inorder[:inorderIndex])+1], inorder[:inorderIndex])
	//右边为右子树的内容，即inorder[inorderIndex+1:]；而对preorder来说，等于上一步中preorder剩下的部分
	root.Right = buildTree(preorder[len(inorder[:inorderIndex])+1:], inorder[inorderIndex+1:])

	return root
}
