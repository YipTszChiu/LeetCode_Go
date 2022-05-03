package _00HotTopic

// 22. 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	// res 存储返回结果
	res := []string{}

	// 通过左括号和右括号的数量进行递归，只有在左括号数量大于等于右括号数量时符合条件
	var dfs func(lRemain int, rRemain int, path string)
	dfs = func(lRemain int, rRemain int, path string) {
		if len(path) == 2*n {
			// 当 path 到达 2n，即 n 对括号，满足条件返回结果
			res = append(res, path)
			return
		}
		// 只要还有左括号可以使用，都可以追加一个左括号
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		// 但使用右括号的条件是：左括号使用数量 > 右括号使用数量，也就是 左括号剩余数量 < 右括号剩余数量
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}
	// 初始状态可以使用n个左右括号，以及path为空串，调用dfs
	dfs(n, n, "")
	return res
}
