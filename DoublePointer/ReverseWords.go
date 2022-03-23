package DoublePointer

//剑指 Offer 58 - I. 翻转单词顺序
//https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	//左右开端
	i, j := 0, len(s)-1
	//清除左右两边空格
	for s[i] == ' ' {
		i++
		//处理空格串
		if i >= len(s) {
			return ""
		}
	}
	for s[j] == ' ' {
		j--
	}
	//反向遍历
	left := i
	i = j
	res := []byte{}

	for i >= left {
		//遍历第一个单词
		for i >= left && s[i] != ' ' {
			i--
		}
		res = append(res, s[i+1:j+1]...)
		//如果未遍历完整个字符串，在当前res后追加空格
		if i > left-1 {
			res = append(res, ' ')
		}
		//跳过单词之间空格
		for i >= left && s[i] == ' ' {
			i--
		}
		j = i
	}
	return string(res)
}
