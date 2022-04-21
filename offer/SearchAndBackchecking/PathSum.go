package SearchAndBackchecking

//剑指 Offer 34. 二叉树中和为某一值的路径
//https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func pathSum(root *TreeNode, target int) (ans [][]int) {
	path := []int{}
	//函数变量
	var dfs func(*TreeNode, int)
	//定义dfs
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		//left减去当前节点val用于判断是否满足路径和为target
		left -= node.Val
		//将当前节点追加到path数组
		path = append(path, node.Val)
		//延迟进行path的回退，注意必须在下面的if之前进行，否则当执行if return时不会进行回退
		defer func() { path = path[:len(path)-1] }()
		//到达叶子节点且和为target（这里为left减到0）
		if node.Left == nil && node.Right == nil && left == 0 {
			//将path复制并追加到ans
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)

	}
	//正式调用dfs
	dfs(root, target)
	return
}
