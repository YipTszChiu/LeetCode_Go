package SearchAndBackchecking

//236. 二叉树的最近公共祖先
//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	//越过叶子节点，或成功找到p或q，则直接返回
	if root == nil || root == q || root == p {
		return root
	}

	//类似后序遍历的思路
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//某一节点的left right均为空，说明没有找到p，q，返回nil
	if left == nil && right == nil {
		return nil
	}
	//如果left空，right不空，则在该节点的右子树中找到p或q，将找到的节点（即right）返回
	if left == nil {
		return right
	}
	//如果left不空，right空，则在该节点则左子树中找到p或q，将找到的节点（即left）返回
	if right == nil {
		return left
	}
	//如果left，right均不为空，说明p，q分别在该节点的左右两边，当前节点即为p，q的祖先
	return root
}
