package Simulation

// 剑指 Offer 31. 栈的压入、弹出序列
// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

func validateStackSequences(pushed []int, popped []int) bool {
	// 思路：按照pushed顺序入栈，如果入栈元素等于popped首元素，则出栈；查看最后剩余的栈情况判断
	// 临时栈temp
	temp := []int{}

	// i用于记录popped的inedx
	i := 0
	// 遍历pushed序列进行入栈
	for _, val := range pushed {
		temp = append(temp, val)
		// 判断栈顶元素与popped序列是否一致，如果是则出栈
		for temp[len(temp)-1] == popped[i] && len(temp) > 0 {
			i++
			temp = temp[:len(temp)-1]
		}
	}

	// 遍历后如果临时栈为空，说明按照popped顺序出栈成功，符合条件
	if len(temp) == 0 {
		return true
	} else {
		return false
	}
}
