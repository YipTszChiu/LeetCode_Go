package _00HotTopic

func combinationSum(candidates []int, target int) [][]int {
	// res 用于返回结果
	res := [][]int{}
	// 定义一个深度搜索函数，temp 为遍历暂存的元素，sum 为 temp 中元素的和
	var dfs func(start int, temp []int, sum int)

	dfs = func(start int, temp []int, sum int) {
		// 遍历的和大于或等于 target
		if sum >= target {
			// 和恰好为 target，该序列满足要求，将该序列加入到 res
			if sum == target {
				newTmp := make([]int, len(temp))
				copy(newTmp, temp)
				res = append(res, newTmp)
			}
			// 和大于 target，不满足条件，直接返回；== 操作结束后也需要返回
			return
		}
		// 遍历的和小于 target，从 start 位遍历 candidates
		for i := start; i < len(candidates); i++ {
			// 将当前遍历的 candidates[i] 加入到 temp 暂存数组
			temp = append(temp, candidates[i])
			// 继续深度搜索，由于 candidates 中的元素可以重复使用，start 仍然从 i 开始遍历
			dfs(i, temp, sum+candidates[i])
			// 删除上面选择的最后一个元素进行剪枝
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0, []int{}, 0)
	return res
}
