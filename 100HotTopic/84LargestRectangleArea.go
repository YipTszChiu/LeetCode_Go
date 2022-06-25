package _00HotTopic

// 84. 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram/

// 方法一：单调栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	// 记录左边 / 右边 最小且最接近的柱子
	left, right := make([]int, n), make([]int, n)
	// 单调栈记录 index
	stack := []int{}

	// 对每个 heights 中的元素进行遍历，记录 heights[i] 左边最小最近的柱子
	for i := 0; i < n; i++ {
		// 栈内记录最小且最接近 i 的柱子
		// 当栈非空且栈顶 index 指向的 heights 元素大于等于当前遍历的 heights[i] 元素，出栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			// 栈内只剩下比 height[i] 小的 index
			stack = stack[:len(stack)-1]
		}
		// 当栈为空，使用一个哨兵记录 height[i] 左边最小最近柱子
		if len(stack) == 0 {
			left[i] = -1
		} else {
			// 栈非空，栈顶 index 指向的 heights 元素即为左边最近最小柱子
			left[i] = stack[len(stack)-1]
		}
		// 当前遍历元素的 index 入栈
		stack = append(stack, i)
	}

	// 清空栈
	stack = []int{}

	// 对每个 heights 中的元素进行遍历，记录 heights[i] 右边最小最近的柱子
	for i := n - 1; i >= 0; i-- {
		// 同理，大于 height[i] 的全部出栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		// 如果栈为空，使用哨兵代替
		if len(stack) == 0 {
			right[i] = n
		} else {
			// 栈非空，栈顶 index 指向的 height 元素即为右边最近最小柱子
			right[i] = stack[len(stack)-1]
		}
		// 当前 index 入栈
		stack = append(stack, i)
	}

	// 计算结果
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}

	return ans
}
