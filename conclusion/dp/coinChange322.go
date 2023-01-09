package dp

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/

func coinChange(coins []int, amount int) int {
	// 零值处理
	if amount == 0 {
		return 0
	}

	// 初始状态
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			// 硬币的面值不大于需要凑的数额
			if coins[j] <= i {
				dp[i] = min322(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min322(i, j int) int {
	if i < j {
		return i
	}
	return j
}
