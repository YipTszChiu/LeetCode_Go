package _00HotTopic

// 309. 最佳买卖股票时机含冷冻期
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/?favorite=2cktkvj

// 动态规划
//思路：使用二维数组进行动态规划, dp[i] 表示第i天结束后的累计最大收益
//dp[i][0] 表示持有股票；dp[i][1] 表示不持有股票并处于冷冻期；dp[i][2] 表示不持有股票且不处于冷冻期
//注意此处处于冷冻期指的是第i天结束之后的状态，第i天结束后处于冷冻期，i+1天无法购买股票
//转移方程1：dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])  第i天结束持有股票：第i-1天本身就持有，或第i天买入
//转移方程2：dp[i][1] = dp[i-1][0] + prices[i] 第i天结束后处于冷冻期，说明当天卖出股票，则第i-1天需要持有股票，并在第i天卖出
//转移方程3：dp[i][2] = max(dp[i-1][1], dp[i-1][2]) 第i天结束不持有股票，且不处于冷冻期，说明当天没有任何操作，则第i-1天不能持有股票
//一共有 n 天，则累计最大收益为 max(dp[n-1][0], dp[n-1][1], dp[n-1][2])
//但最后一天仍持有股票的话没有意义，因此只需要考虑 max(dp[n-1][1], dp[n-1][2])
func maxProfit_309(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	length := len(prices)
	dp := make([][3]int, length)
	// 初始状态：第0天如果持有股票，只能在第0天买入，当天结束收益dp[0][0] = -prices[0]
	dp[0][0] = -prices[0]
	for i := 1; i < length; i++ {
		// 注意此处只用到了前一天的数据，因此只需要用同一个变量存储，可以进行空间优化
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[length-1][1], dp[length-1][2])
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
