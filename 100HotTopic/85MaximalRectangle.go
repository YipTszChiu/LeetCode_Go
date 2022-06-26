package _00HotTopic

// 85. 最大矩形
// https://leetcode.cn/problems/maximal-rectangle/

// 方法一：暴力破解的优化
func maximalRectangle1(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	// 记录矩阵的行数和列数分别为 m, n
	m, n := len(matrix), len(matrix[0])
	// left 矩阵记录每个 matrix[i][j] 左边（包括自己）有多少个 "1"
	left := make([][]int, m)

	// 按行遍历 matrix
	for i, row := range matrix {
		// 初始化 left
		left[i] = make([]int, n)
		// 按列遍历 matrix[i]
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	// 遍历 matrix，以 matrix[i][j] 为矩形的右下角，计算最大矩形
	for i, row := range matrix {
		for j, v := range row {
			// 0 无法构成矩形，直接跳过
			if v == '0' {
				continue
			}
			// 记录宽和面积
			width := left[i][j]
			area := width
			// 以 matrix[i][j] 为右下角， i-- 向上遍历
			for k := i - 1; k >= 0; k-- {
				// 宽度取最小 left
				width = min(width, left[k][j])
				// 计算面积
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}

	return
}

// 方法二：结合单调栈
func maximalRectangle2(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	// m, n 存储 matrix 矩阵的行列数
	m, n := len(matrix), len(matrix[0])
	// left 矩阵记录每个点左边有多少个 "1"
	left := make([][]int, m)
	// 按行遍历 matrix
	for i, row := range matrix {
		// 顺便初始化 left 矩阵，每行为 n 列
		left[i] = make([]int, n)
		// 按列遍历 matrix[i]
		for j, v := range row {
			// "0" 无法构成矩形，跳过
			if v == '0' {
				continue
			}
			// 首列只有自己，因此 left = 1
			if j == 0 {
				left[i][j] = 1
			} else {
				// 其他非0列 left 值 + 1
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		// 以列为基础，对每一列从上到下看，使用 84 的单调栈
		for i, l := range left {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
				stk = stk[:len(stk)-1]
			}
			up[i] = -1
			if len(stk) > 0 {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = nil
		// 以列为基础，对每一列从下到上看，使用 84 的单调栈
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			down[i] = m
			if len(stk) > 0 {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		// 以列看做宽，left 看做高，计算矩形面积
		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			ans = max(ans, area)
		}
	}
	return
}
