package _00HotTopic

// 20. 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	// 使用 slice 构造一个 stack 用于括号匹配
	stack := []byte{}
	// 按照 s 的顺序入栈
	for i := 0; i < len(s); i++ {
		// 如果是左括号则直接入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 {
			// 首先判断 stack 长度是否大于零，在未遍历完 s 出现空栈说明：在没有左括号的情况下输入了一个右括号
			// 如果是右括号则与栈顶进行匹配，如果成功则栈顶元素直接出栈
			switch stack[len(stack)-1] {
			case '(':
				if s[i] == ')' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case '[':
				if s[i] == ']' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case '{':
				if s[i] == '}' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		} else {
			// 在没有左括号的情况下输入了一个右括号
			return false
		}
	}
	// 遍历后如果 stack 为空，说明全部匹配成功
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
