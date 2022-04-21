package DynamicProgramming

//剑指 Offer 47. 礼物的最大价值
//https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

//只能往右或者下，即每个点只能从上边或者左边到达
func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			//起始点无需处理
			if i == 0 && j == 0 {
				continue
			}
			//第一行除起始点以外的元素，只能从左边到达
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				//第一列除了起始点以外的元素，只能从上边到达
				grid[i][j] += grid[i-1][j]
			} else {
				//除了第一行和第一列的特例，其余点均可以从左边或者上边到达
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	//返回终点的值
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
