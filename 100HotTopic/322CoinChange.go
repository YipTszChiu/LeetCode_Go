package _00HotTopic

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/?favorite=2cktkvj

// 动态规划
// 时间复杂度：O(Sn) S代表总金额，n代表硬币数
// 空间复杂度：O(S) S代表总金额
func coinChange(coins []int, amount int) int {
	// dp[i] 表示组成金额 i 所需最少硬币数， dp[i] = min(dp[i-cj] + 1)，即最后使用一枚面值为cj的硬币可以凑出i
	dp := make([]int, amount+1)
	// 初始化dp数组，由于后面需要用min比较，因此初始化一个不可能组成的数amount+1
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	// 初始状态 dp[0] = 0，因为面额0无法用硬币凑出
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
