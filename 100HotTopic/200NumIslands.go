package _00HotTopic

// 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands/?favorite=2cktkvj

// 深度优先遍历 dfs
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || grid == nil {
		return 0
	}
	row := len(grid)
	column := len(grid[0])

	// 思路：深度遍历整个 grid, 当碰到 1 后将当前位置 0, 岛屿数+1
	var dfs func(grid [][]byte, row, column int)
	dfs = func(grid [][]byte, r, c int) {
		// 越过 grid 边界直接返回 || 如果当前位置本身为 0 也返回
		if r < 0 || c < 0 || r >= row || c >= column || grid[r][c] == '0' {
			return
		}
		// 当前位置为 1 则将其置 0 并继续深度遍历
		grid[r][c] = '0'
		dfs(grid, r+1, c)
		dfs(grid, r-1, c)
		dfs(grid, r, c+1)
		dfs(grid, r, c-1)
	}
	// count 用于记录岛屿数
	count := 0
	// 遍历 grid
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}
