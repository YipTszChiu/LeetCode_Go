package _00HotTopic

// 240. 搜索二维矩阵 II
// https://leetcode.cn/problems/search-a-2d-matrix-ii/?favorite=2cktkvj

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func searchMatrix(matrix [][]int, target int) bool {
	row, column := len(matrix), len(matrix[0])
	i, j := 0, column-1
	// 从右上角遍历整个矩阵
	for i < row && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			j++
		} else {
			i--
		}
	}
	return false
}
