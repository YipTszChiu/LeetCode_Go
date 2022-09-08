package _00HotTopic

// 221.最大正方形
// https://leetcode.cn/problems/maximal-square/solution/zui-da-zheng-fang-xing-by-leetcode-solution/

// 动态规划
func maximalSquare(matrix [][]byte) int {
	// dp[i][j] 代表以 matrix[i][j] 为右下角作为正方形的边长最大值
	// 初始化dp矩阵
	dp := make([][]int, len(matrix))
	maxSide := 0
	// 将所有matrix = 1 位置的 dp 初始值置为 1
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j]) - '0'
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	// 开始进行dp计算
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

// 暴力破解
func maximalSuqare(matrix [][]byte) int {
	maxSide := 0
	// 零值处理
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}
	rows, columns := len(matrix), len(matrix[0])
	// 遍历 matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				// 以该 '1' 为左上角，将其右下方都视为正方形
				curMaxSide := min(rows-i, columns-j)
				// 遍历右下方是否为 '1'
				for k := 1; k < curMaxSide; k++ {
					flag := true
					// 检测右下角
					if matrix[i+k][j+k] == '0' {
						break
					}
					// 使用 m 可以避开重复检测，每次只检测最外层
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}
