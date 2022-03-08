package SearchAndBackchecking

//面试题32 - I. 从上到下打印二叉树
//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

func levelOrder(root *TreeNode) []int {
	//用于保存返回值
	ans := []int{}
	//用于层次遍历二叉树的队列
	que := []*TreeNode{}

	if root == nil {
		return nil
	}

	//将root入队
	que = append(que, root)
	//开始层次遍历
	for len(que) > 0 {
		if que[0].Left != nil {
			que = append(que, que[0].Left)
		}
		if que[0].Right != nil {
			que = append(que, que[0].Right)
		}
		ans = append(ans, que[0].Val)
		que = que[1:]
	}

	return ans
}

//剑指 Offer 32 - II. 从上到下打印二叉树 II
//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	//ans用于返回结果
	ans := [][]int{}
	//temp用于暂存每一层的元素
	temp := []int{}
	//que是用于层次遍历的队列，初始化将root入队
	que := []*TreeNode{root}
	//设置两个指针，ptr指向当前遍历位置，levelPtr指向每一层最右边的节点
	ptr, levelPtr := 0, 0

	//开始层次遍历
	for ptr < len(que) {
		//左右子树入队
		if que[ptr].Left != nil {
			que = append(que, que[ptr].Left)
		}
		if que[ptr].Right != nil {
			que = append(que, que[ptr].Right)
		}

		//当前节点保存到temp
		temp = append(temp, que[ptr].Val)

		//如果ptr = levelPtr，证明遍历到该层最后一个节点
		if ptr == levelPtr {
			//更新levelPtr
			levelPtr = len(que) - 1
			ans = append(ans, temp)
			//清空切片
			temp = []int{}
		}

		//当前节点出队
		ptr++

	}

	return ans
}

//剑指 Offer 32 - III. 从上到下打印二叉树 III
//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	//ans用于返回结果，temp用于暂存每一层的元素
	ans := [][]int{}
	temp := []int{}

	//que用于层次遍历，初始化将root入队
	que := []*TreeNode{root}

	//ptr指向当前节点，levelPtr指向每一层最后一个节点，初始化时都为0
	//tag用于判断奇偶，判断是否需要反转
	ptr, levelPtr, tag := 0, 0, 0

	//开始层次遍历
	for ptr < len(que) {
		//左右子节点入队
		if que[ptr].Left != nil {
			que = append(que, que[ptr].Left)
		}
		if que[ptr].Right != nil {
			que = append(que, que[ptr].Right)
		}

		//当前节点出队
		temp = append(temp, que[ptr].Val)

		//如果当前节点为该层最后一个节点，更新levelPtr，向ans中追加并重置temp，tag加1
		if ptr == levelPtr {
			//无需反转
			if tag%2 == 0 {
				ans = append(ans, temp)
			} else {
				ans = append(ans, reverseArr(temp))
			}
			levelPtr = len(que) - 1
			temp = []int{}
			tag++
		}
		//当前节点出队完成，ptr指向队内下一个节点
		ptr++
	}

	return ans
}

func reverseArr(arr []int) []int {
	temp := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		temp = append(temp, arr[i])
	}

	return temp
}
