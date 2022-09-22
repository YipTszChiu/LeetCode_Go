package _00HotTopic

// 301. 删除无效的括号
// https://leetcode.cn/problems/remove-invalid-parentheses/?favorite=2cktkvj

// isValid 判断一个字符串的左右括号是否匹配
func isStrValid(str string) bool {
	count := 0
	// 遍历字符串
	for _, c := range str {
		// 遇到左括号计数值加一
		if c == '(' {
			count++
			// 遇到右括号计数值减一
		} else if c == ')' {
			count--
			// 如果计数值小于零，即出现无法匹配的右括号，则括号匹配失败，返回false
			if count < 0 {
				return false
			}
		}
	}
	// 循环结束判断 count 是否恰好为 0，如果大于 0 则出现无法匹配的左括号
	return count == 0
}

func helper(ans *[]string, str string, start, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		// 字符串括号匹配成功，作为结果加入 ans
		if isStrValid(str) {
			*ans = append(*ans, str)
		}
		return
	}
	// 从 start 位置开始遍历 str
	for i := start; i < len(str); i++ {
		// 遇到连续重复的括号，进行剪枝
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 剩余字符已经不足以完成删除，直接返回
		if lremove+rremove > len(str)-i {
			return
		}
		// 尝试删除一个左括号
		if lremove > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, lremove-1, rremove)
		}
		// 尝试删除一个右括号
		if rremove > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, lremove, rremove-1)
		}
	}
}

func removeInvalidParentheses(s string) (ans []string) {
	// 统计需要删除的左右括号数量
	lremove, rremove := 0, 0
	for _, c := range s {
		if c == '(' {
			lremove++
		} else if c == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}
	// 调用 helper 进行递归尝试删除括号
	helper(&ans, s, 0, lremove, rremove)
	return
}
