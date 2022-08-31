package _00HotTopic

// 198. 打家劫舍
//	https://leetcode.cn/problems/house-robber/?favorite=2cktkvj

// 动态规划
// 时间复杂度：O(n) 遍历一次数组
// 空间复杂度：O(n) 使用一个数组存储dp
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// 思路：
	// dp[i] 代表前 i 间屋子能偷窃到的最高金额，转移方程如下：
	// dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	// 说明：相邻屋子不能连续偷取，因此两个选择为偷取i-1不偷取i，或偷取i-2以及i
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

// 空间优化：由于dp只与i-1与i-2有关，因此使用 first 与 second 存储
// 时间复杂度：O(n) 遍历一次数组
// 空间复杂度：O(1) 使用常数空间
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
