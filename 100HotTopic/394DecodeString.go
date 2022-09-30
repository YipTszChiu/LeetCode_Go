package _00HotTopic

import (
	"strconv"
	"strings"
)

// 394. 字符串解码
// https://leetcode.cn/problems/decode-string/?favorite=2cktkvj

// 栈
func decodeString(s string) string {
	stack := []string{}
	ptr := 0
	for ptr < len(s) {
		// 获取当前遍历的字符
		cur := s[ptr]
		// 判断该字符
		if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || (cur == '[') {
			// 当前字符为字母或左括号，直接入栈
			stack = append(stack, string(cur))
			ptr++
		} else if cur >= '0' && cur <= '9' {
			// 当前字符串为数字，需要将单独的一个字符拼接为一个字符串
			digits := ""
			for ; s[ptr] >= '0' && s[ptr] <= '9'; ptr++ {
				digits += string(s[ptr])
			}
			stack = append(stack, digits)
		} else {
			// 当前字符为右括号，出栈并处理左右括号之间的字符
			ptr++
			subStr := []string{}
			// 一直出栈到左括号为止，并将出栈的字符存到subStr中
			for stack[len(stack)-1] != "[" {
				// 此处是倒序的字符串
				subStr = append(subStr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 翻转subStr
			for i := 0; i < len(subStr)/2; i++ {
				subStr[i], subStr[len(subStr)-1-i] = subStr[len(subStr)-1-i], subStr[i]
			}
			// 将左括号出栈
			stack = stack[:len(stack)-1]
			// 左括号左边为数字，将其转换为int
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			// 将该数字出栈
			stack = stack[:len(stack)-1]
			t := strings.Repeat(toString(subStr), repTime)
			stack = append(stack, t)
		}
	}
	return toString(stack)
}

func toString(str []string) string {
	s := ""
	for i := 0; i < len(str); i++ {
		s += str[i]
	}
	return s
}
