package _00HotTopic

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum/?favorite=2cktkvj

// 递减队列
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	// push 维护递减队列
	push := func(index int) {
		for len(queue) > 0 && nums[index] >= nums[queue[len(queue)-1]] {
			// 将所有小于等于 val 的值出队
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, index)
	}

	// 初始化递减队列
	for i := 0; i < k; i++ {
		push(i)
	}

	// result 为滑动窗口的最大值序列
	result := []int{nums[queue[0]]}

	// 对剩余的元素进行遍历
	for i := k; i < len(nums); i++ {
		// 将最新的值的索引入队
		push(i)
		// 如果最大值的 index 已经不在窗口内，将该 index 出队
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		// result 记录当前最大值
		result = append(result, nums[queue[0]])
	}
	return result
}
