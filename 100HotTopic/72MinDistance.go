package _00HotTopic

// 72. 编辑距离
// https://leetcode.cn/problems/edit-distance/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// 如果有一个为空串，返回非空串的长度
	if m == 0 || n == 0 {
		return m + n
	}

	// 初始化dp数组
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// 空串与非空串编辑距离为非空串长度
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	// 三个操作有等效的，最终总结只有三种不同操作：
	// a. 给word1增加一个字符；
	// b. 给word2增加一个字符；
	// c. word1改变一个字符；
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			a := dp[i-1][j] + 1
			b := dp[i][j-1] + 1
			c := dp[i-1][j-1]
			// 判断 i-1, j-1 是否为相同字符，如果不是需要 c+1, 如果相同则不需要 +1
			if word1[i-1] != word2[j-1] {
				c += 1
			}
			// 取最小作为编辑长度
			dp[i][j] = min(a, min(b, c))
		}
	}

	return dp[m][n]
}
