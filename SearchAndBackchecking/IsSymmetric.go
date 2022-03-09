package SearchAndBackchecking

//剑指 Offer 28. 对称的二叉树
//https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/

func isSymmetric(root *TreeNode) bool {
	//空值处理
	if root == nil {
		return true
	}

	//root根节点本身不用判断是否对称
	//调用judge判断root的左右子树是否对称
	return judge(root.Left, root.Right)
}

func judge(left *TreeNode, right *TreeNode) bool {
	//同时为空也对称
	if left == nil && right == nil {
		return true
	}
	//Val不等，或者只有一边为空，不对称
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	//注意：由于是对称，必须是镜像对比
	return judge(left.Left, right.Right) && judge(left.Right, right.Left)
}
