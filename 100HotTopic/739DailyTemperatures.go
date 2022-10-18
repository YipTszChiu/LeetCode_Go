package _00HotTopic

// 739. 每日温度
// https://leetcode.cn/problems/daily-temperatures/?favorite=2cktkvj

// 暴力破解超时
func dailyTemperatures(temperatures []int) []int {
	ans := []int{}
	// 固定一天开始遍历
	for i := 0; i < len(temperatures); i++ {
		count, j := 0, i+1
		// 遍历该天后第一个比他高温的日子
		for ; j < len(temperatures); j++ {
			count++
			if temperatures[i] < temperatures[j] {
				break
			}
		}
		// 如果j没有越界，说明有出现高温，计数值即为高温距离天数
		if j < len(temperatures) {
			ans = append(ans, count)
		} else {
			// j越界说明不存在高温，追加0
			ans = append(ans, 0)
		}
	}
	return ans
}

// 单调栈
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	// 维护一个存储下标的栈
	stack := []int{}
	// ans 存储返回结果，使用make构造，使其默认值为 0
	ans := make([]int, n)
	// 顺序遍历数组
	for i := 0; i < n; i++ {
		// 记录当前的温度值
		temperature := temperatures[i]
		// 对后续日期进行遍历：1. 栈非空，2.当前温度大于栈顶index那天的温度（说明当前温度是栈顶那天第一个高温）
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			// 修改返回结果，栈顶日期索引为：stack[len(stack)-1]，当前日期索引为 i，相减获得日期差
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// 栈顶索引日期已经找到比他高温的日期，出栈
			stack = stack[:len(stack)-1]
		}
		// 当前日期索引入栈
		stack = append(stack, i)
	}
	return ans
}
