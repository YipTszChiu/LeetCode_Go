package _00HotTopic

// 437. 路径总和 III
// https://leetcode.cn/problems/path-sum-iii/?favorite=2cktkvj

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 深度优先遍历
// 时间复杂度：O(N^2) 需要以每个节点为根进行遍历，最大花费时间为O(N)
// 空间复杂度：O(N) 栈空间
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

// 前缀和
// 辅助函数，用于前序遍历二叉树，同时记录下遍历过程中的前进路径和以及出现的次数
// 输入是二叉树，前进路径和，目标值sum，以及辅助哈希表,输出是路径和等于sum的路径数量
func dfs(root *TreeNode, cur, sum int, prefixSum map[int]int) int {
	if root == nil { // 如果树为空
		return 0 // 就返回数量0
	}
	cur += root.Val // 否则前进路径和加上当前节点值进行更新
	// 然后计算前进路径和与目标值sum的差值,并在哈希中查找是否存在
	cnt := 0 // 用0初始化路径数量
	if v, ok := prefixSum[cur-sum]; ok {
		cnt = v // 存在的话就取出对应值，并初始化路径数量，
	}
	// 更新前进路径和cur出现的次数
	prefixSum[cur]++ // 如果是第一次出现则会设置为1
	// 接下来递归求解左右子树,并把返回的路径数量加到cnt上
	cnt += dfs(root.Left, cur, sum, prefixSum)
	cnt += dfs(root.Right, cur, sum, prefixSum)
	prefixSum[cur]-- // 退递归后需要把现在的前进路径和数量-1
	// 下一步要回溯到递归的上一层进行处理
	return cnt // 最后返回路径数量即可
}

// 保存前进路径和的方法 Time: O(n), Space: O(n)
func pathSum2(root *TreeNode, sum int) int {
	prefixSum := make(map[int]int)      // 定义辅助哈希表用于存储每个前进路径和以及它出现的次数
	prefixSum[0] = 1                    // 接着把处理边界的0，1数对加入到哈希表
	return dfs(root, 0, sum, prefixSum) // 最后只要调用辅助函数即可
}

//作者：CusterFun
//链接：https://leetcode.cn/problems/path-sum-iii/solution/custerxue-xi-bi-ji-er-cha-shu-de-di-gui-he-dfs-by-/
