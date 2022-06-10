package _00HotTopic

// 62.不同路径
// https://leetcode.cn/problems/unique-paths/

// dp
func uniquePaths(m, n int) int {
	// 初始化二维矩阵
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		// 第一列均可到达
		dp[i][0] = 1
	}
	// 第一行均可到达
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 每个点只能从左边或者上边到达
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//// 自己使用 dfs 超时
//var count int
//
//func uniquePaths(m int, n int) int {
//	count = 0
//	dfs62(0, 0, m, n)
//	return count
//}
//
//func dfs62(i, j, m, n int) {
//	// 如果到达终点则返回成功
//	if i == m-1 && j == n-1 {
//		count++
//	}
//	// 越界则返回失败
//	if i >= m || j >= n {
//		return
//	}
//	// 正常情况则向下或向右遍历
//	dfs62(i+1, j, m, n)
//	dfs62(i, j+1, m, n)
//}
