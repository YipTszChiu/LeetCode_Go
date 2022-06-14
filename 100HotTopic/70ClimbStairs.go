package _00HotTopic

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/

func climbStairs(n int) int {
	dp := make([]int, 3)
	// 一级，二级台阶的方法都只有一种
	dp[1], dp[2] = 1, 2
	if n <= 2 {
		return dp[n]
	}

	// 遍历
	for i := 3; i <= n; i++ {
		dp[0] = dp[1]
		dp[1] = dp[2]
		dp[2] = dp[0] + dp[1]
	}

	return dp[2]
}
