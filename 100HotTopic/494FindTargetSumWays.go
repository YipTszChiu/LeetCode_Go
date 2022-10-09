package _00HotTopic

// 494. 目标和
// https://leetcode.cn/problems/target-sum/?favorite=2cktkvj

// 回溯
// 时间复杂度：O(2^n) 每个数都可以选择正或者负
// 空间复杂度：O(n) 递归栈的深度
func findTargetSumWays(nums []int, target int) int {
	// count 记录返回的结果
	count := 0
	// 类似二叉树深度遍历
	var dfs func(index, sum int)
	dfs = func(index, sum int) {
		// 遍历完整个数组，判断当前和是否为target
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		// 每一个数都可以加或减，类似二叉树左右子树
		sumPos := sum + nums[index]
		sumNeg := sum - nums[index]
		// 递归并回溯
		dfs(index+1, sumPos)
		dfs(index+1, sumNeg)
	}
	dfs(0, 0)
	return count
}
