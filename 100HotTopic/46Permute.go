package _00HotTopic

// 46. 全排列
// https://leetcode.cn/problems/permutations/

func permute(nums []int) [][]int {
	// res 用于返回所有的全排列
	res := [][]int{}
	n := len(nums)

	// 回溯算法：tag 为 nums 数组中每个数字是否使用， path 记录当前排列
	var backtrack2 func(tag []bool, path []int)
	backtrack2 = func(tag []bool, path []int) {
		// 完成遍历，将结果添加到 res
		if len(path) == n {
			// 复制问题
			temp := make([]int, n)
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			// 如果当前数字已经使用则跳过
			if tag[i] == true {
				continue
			}
			// 否则使用当前数字
			tag[i] = true
			path = append(path, nums[i])
			// 继续进行遍历
			backtrack2(tag, path)
			// 回溯
			path = path[:len(path)-1]
			tag[i] = false
		}
	}

	backtrack2(make([]bool, n), []int{})

	return res

}
