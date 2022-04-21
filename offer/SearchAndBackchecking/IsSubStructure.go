package SearchAndBackchecking

//剑指 Offer 26. 树的子结构
//https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	//空值处理
	if A == nil || B == nil {
		return false
	}

	//先序遍历
	//先调用recur判断B是否为A的子结构，然后判断B是否为A.Left，A.Right的子结构
	//注意：isSubStructure的功能相当于先序遍历，recur相当于先序遍历中对根的操作
	//因此不是return recur(A, B) || recur (A.Left, B) || recur (A.Right, B)
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A *TreeNode, B *TreeNode) bool {
	//越过B叶子节点，即对B每个节点已经满足A.Val = B.Val，返回true
	if B == nil {
		return true
	}
	//越过A叶子节点，即B的深度比A大，B一定不是A的子结构，返回false
	if A == nil {
		return false
	}
	//对应节点Val不等，返回false
	if A.Val != B.Val {
		return false
	}

	//如果A.Val = B.Val，继续判断A，B的左右子树是否满足
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
