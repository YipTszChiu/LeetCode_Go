package _00HotTopic

// 32. 最长有效括号
// https://leetcode-cn.com/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	// 初始化最大计数值
	maxAns := 0
	// 初始化一个栈底为“最后一个没有被匹配的右括号的下标”的栈
	// 为了避免一开始是左括号直接进空栈，与上面目标不符，因此初始化栈时有一个 -1 的元素
	stack := []int{-1}
	// 遍历字符串 s
	for i := 0; i < len(s); i++ {
		// 碰到左括号则入栈
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 碰到右括号则将栈顶元素出栈
			stack = stack[:len(stack)-1]
			// 出栈后栈为空，需要将当前这个“最后一个没有被匹配的右括号”的下标入栈
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				// 出栈后不为空，说明括号匹配成功，当前索引剪去栈顶等于总共匹配成功的括号数值
				// 说明：如果一直输入右括号，栈底与栈顶指向同一个位置，会不断更新最后一个右括号的索引；
				//      如果存在左括号，此时与右括号匹配成功已经出栈，此时栈顶为 最后一个右 or 最后一个左（连续左的情况）
				// 		无论是哪种情况，当前的索引减去栈顶即为长度，用这个长度去和 maxAns 比较大小
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}

	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
