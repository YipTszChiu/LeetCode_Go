package _00HotTopic

// 121. 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

// 动态规划
func maxProfit(prices []int) int {
	// 第 i 天的利润只与之前的最低价格有关
	// dp[i] 表示第 i 天及之前最低价格
	max := 0
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		// dp[i] = min(dp[i-1], prices[i])
		if prices[i] < dp[i-1] {
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i-1]
		}
		// max = max(prices[i] - dp[i], max)
		if prices[i]-dp[i] > max {
			max = prices[i] - dp[i]
		}
	}

	return max
}

// 优化：观察到 dp[i] 仅在 prices[i] < dp[i-1] 时改变，因此只需要用一个变量 minPrice 记录
func maxProfit2(prices []int) int {
	// 第 i 天的利润只与之前的最低价格有关
	// dp[i] 表示第 i 天及之前最低价格
	max := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > max {
			// 此处也有优化：当 prices[i] 更新时（即上面的if成立），此时prices[i] - minPrice = 0，max无需更新
			// 因此仅有上面的 if 不成立（即prices[i] >= minPrice），且 prices[i] - minPrice > max 时，max需要更新
			max = prices[i] - minPrice
		}
	}

	return max
}
