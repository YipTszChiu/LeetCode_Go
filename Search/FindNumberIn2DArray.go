package Search

//二维数组中的查找
//https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	//从左下角开始查找
	row := len(matrix)-1
	column := 0

	for row >= 0 && column <= len(matrix[0])-1 {
		if matrix[row][column] == target {
			//找到target直接返回
			return true
		} else if matrix[row][column] < target {
			//当前数值 < target，说明target一定不会在当前位置上方（从上到下递增）
			column++
		} else {
			//当前数值 > target，说明targe一定不会在当前位置右方（从左到右递增）
			row--
		}
	}
	//遍历二位数组后没有找到target，返回false
	return false
}