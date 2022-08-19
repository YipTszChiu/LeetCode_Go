package _00HotTopic

// 139.单词拆分
// https://leetcode.cn/problems/word-break/

// 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func wordBreak(s string, wordDict []string) bool {
	// 使用 map 映射字典中的字符串是否存在
	wordDictSet := make(map[string]bool)
	// 遍历 wordDict 对 map 进行初始化
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	// dp[i] 表示字符串 s[0..i-1] 是否能被空格拆分成字典里的单词
	dp := make([]bool, len(s)+1)
	// dp[0] 表示空串，设置为合法
	dp[0] = true
	// dp[i] 是否合法，可以从 dp[0..i-1]判断
	// 如 dp[j]合法，那判断 s[j..i-1]能否被字典表示就可知道 dp[i]
	// dp[i] = dp[j] & check(s[j:i-1])
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
