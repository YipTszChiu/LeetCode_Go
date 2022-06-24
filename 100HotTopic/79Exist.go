package _00HotTopic

// 79. 单词搜索 / 剑指offer12
// https://xd2r8gb0a1.feishu.cn/docs/doccnV7SizUChKHwxjtbKWJ35Zb#MCoj6l

func exist(board [][]byte, word string) bool {
	// 记录矩阵的行和列数
	m := len(board)
	n := len(board[0])
	// 使用一个同样大小的矩阵记录是否被遍历
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	// 声明并定义 dfs, i, j, k 分别为正在遍历 board[i][j], word[k]
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 剪枝：遇到不匹配的直接返回fasle
		if board[i][j] != word[k] || i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		// 已经遍历完 word，说明匹配成功，返回true
		if k == len(word)-1 {
			return true
		}
		// 正常遍历
		visit[i][j] = true
		down := i+1 < m && !visit[i+1][j] && dfs(i+1, j, k+1)
		right := j+1 < n && !visit[i][j+1] && dfs(i, j+1, k+1)
		up := i-1 >= 0 && !visit[i-1][j] && dfs(i-1, j, k+1)
		left := j-1 >= 0 && !visit[i][j-1] && dfs(i, j-1, k+1)
		res := down || right || up || left
		// 复原当前位
		visit[i][j] = false
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
