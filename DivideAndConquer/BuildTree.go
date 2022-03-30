package DivideAndConquer

//剑指 Offer 07. 重建二叉树
//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归法
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

//迭代法
func buildTree2(preorder []int, inorder []int) *TreeNode {
	//空值处理
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//构造树根
	root := &TreeNode{
		preorder[0],
		nil,
		nil,
	}

	//构造一个存储TreeNode的栈，并将树根入栈
	stack := []*TreeNode{root}
	//用于中序遍历的指针
	inorderIndex := 0

	for i := 1; i < len(preorder); i++ {
		//记录当前先序遍历的值
		preorderVal := preorder[i]
		//获取栈顶的元素
		node := stack[len(stack)-1]
		//如果栈顶元素不等于中序遍历指针指向的值
		if node.Val != inorder[inorderIndex] {
			//先序遍历的值为栈顶元素的左节点
			node.Left = &TreeNode{preorderVal, nil, nil}
			//新建的节点入栈
			stack = append(stack, node.Left)
		} else {
			//如果栈顶元素等于中序遍历指针指向的值（注意溢出）
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				//栈顶出栈
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				//指针后移
				inorderIndex++
			}
			//找到出栈节点不等于中序遍历指针指向值的一位，preorderVal即为这个节点的右节点
			node.Right = &TreeNode{preorderVal, nil, nil}
			//新建的节点入栈
			stack = append(stack, node.Right)
		}
	}
	return root
}
