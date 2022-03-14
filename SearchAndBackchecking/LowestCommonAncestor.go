package SearchAndBackchecking

//剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/

//解法一：两次遍历
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

func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	pLen, qLen := len(pathP), len(pathQ)
	for i := 0; i < pLen && i < qLen && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i]
	}
	return
}

//解法二：一次遍历
func lowestCommonAncestor2(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root

	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}
