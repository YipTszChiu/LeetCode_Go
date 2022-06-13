package _00HotTopic

// 64. 最小路径和
// https://leetcode.cn/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	// 特值处理
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// 初始化 dp矩阵
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	// 顶部一行只能通过左边到达
	for j := 1; j < columns; j++ {
		dp[0][j] += dp[0][j-1] + grid[0][j]
	}
	// 左边一列只能通过上边到达
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 遍历整个矩阵
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][columns-1]
}
