package Simulation

// 剑指 Offer 29. 顺时针打印矩阵
// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

func spiralOrder(matrix [][]int) []int {
	// 空值处理
	if len(matrix) == 0 {
		return []int{}
	}

	// 初始化四个坐标
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	// res 存储返回结果
	res := []int{}

	// 对整个矩阵进行顺时针遍历
	for {
		// 从左上到右上
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		// top加一，判断top越过bottom
		top++
		if top > bottom {
			break
		}
		// 从右上到右下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		// right减一，判断right是否越过left
		right--
		if right < left {
			break
		}
		// 从右下到左下
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		// bottom减一，判断bottom是否越过top
		bottom--
		if bottom < top {
			break
		}
		// 从左下到左上
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		// left加一，判断left是否越过right
		left++
		if left > right {
			break
		}
	}
	return res
}
