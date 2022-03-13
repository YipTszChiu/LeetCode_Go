package SearchAndBackchecking

//剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/

func getPath(root, target *TreeNode) (path []*TreeNode) {
	node := root
	//遍历直到找到target节点，且题中target必然存在于树中
	for node != target {
		path = append(path, node)
		//二叉搜索树的性质
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	//当node == target时退出循环，将target也存入到path
	path = append(path, node)
	return
}
