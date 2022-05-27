package _00HotTopic

// 48. 旋转图像
// https://leetcode.cn/problems/rotate-image/

func rotate(matrix [][]int) {
	// 先水平翻转 matrix[i] <-> matrix[n-1-i]
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 然后根据主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
