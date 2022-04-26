package _00HotTopic

// 10. 正则表达式匹配
// https://leetcode-cn.com/problems/regular-expression-matching/solution/zheng-ze-biao-da-shi-pi-pei-by-leetcode-solution/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// dp[i][j] 表示 s[0:i] 与 p[0:j] 匹配
	// 初始化一个 [m+1][n+1] 的 bool 矩阵，需要考虑空串和空正则
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	for i := 0; i < m+1; i++ {
		// 非空串和空正则必然不匹配，因此 dp[i][0] 为默认的 false 无需遍历
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
