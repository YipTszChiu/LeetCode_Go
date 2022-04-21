package SearchAndBackchecking

//剑指 Offer 27. 二叉树的镜像
//https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

func mirrorTree(root *TreeNode) *TreeNode {
	//空值处理
	if root == nil {
		return nil
	}
	//先序遍历，对每个root交换其左右子树
	left := root.Left
	right := root.Right
	root.Left = right
	root.Right = left
	//遍历左右子树
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	//返回根节点
	return root

}
