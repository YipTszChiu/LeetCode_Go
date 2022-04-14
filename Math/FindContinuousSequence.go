package Math

// 剑指Offer 57 - II. 和为s的连续正数序列
// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

func findContinuousSequence(target int) [][]int {
	// 构建用于返回的二维数组
	res := [][]int{}
	// 滑动窗口初始化
	left, right, sum := 1, 2, 3
	// 进行滑动遍历
	for left < right {
		// 找到一组连续正数符合sum = target
		if sum == target {
			// 用于返回当前窗口的元素
			temp := []int{}
			// 将当前窗口中的元素追加到temp
			for i := left; i <= right; i++ {
				temp = append(temp, i)
			}
			// 将该序列追加到返回的结果
			res = append(res, temp)
		}
		// 窗口中sum小于target，右指针往右移
		if sum < target {
			right++
			sum += right
		} else {
			// 窗口中sum大于target，左指针往右移
			sum -= left
			left++
		}
	}

	return res
}
