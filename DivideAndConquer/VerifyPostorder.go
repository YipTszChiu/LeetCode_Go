package DivideAndConquer

//剑指 Offer 33. 二叉搜索树的后序遍历序列
//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

//递归法
func verifyPostorder(postorder []int) bool {
	//基本思路：二叉搜索树满足 左 < 根 < 右
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}

	//后序遍历的末尾即为根节点的值
	rootVal := postorder[end]
	//找到第一个大于根节点值的元素，即为右子树中的元素
	i := start
	for ; i < end; i++ {
		if postorder[i] > rootVal {
			break
		}
	}
	//如果右子树中的元素小于根节点，不符合二叉搜索树的性质
	for j := i; j < end; j++ {
		if postorder[j] < rootVal {
			return false
		}
	}

	//递归判断左右子树是否满足二叉搜索树性质
	return recur(postorder, start, i-1) && recur(postorder, i, end-1)
}

//非递归法
func verifyPostorder2(postorder []int) bool {
	//如果只有两个节点无论如何都能构成二叉搜索树
	if len(postorder) <= 2 {
		return true
	}

	//后序遍历的末尾即为根节点
	root := len(postorder) - 1
	//遍历所有节点查看是否满足左 < 根 < 右
	for root > 0 {
		pointer := 0
		//左子树小于根节点
		for postorder[pointer] < postorder[root] {
			pointer++
		}
		//右子树大于根节点
		for postorder[pointer] > postorder[root] {
			pointer++
		}
		//经过遍历后，如果满足 左 < 根 < 右，pointer应恰好指向根节点
		if pointer != root {
			return false
		}

		//所有节点按照这个规则遍历一遍：如果一棵树是二叉搜索树，那么他的子树也是二叉搜索树
		root--
	}

	return true
}
