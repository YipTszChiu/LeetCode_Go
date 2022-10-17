package _00HotTopic

// 647. 回文子串
// https://leetcode.cn/problems/palindromic-substrings/

// 中心拓展
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func countSubstrings(s string) int {
	ans := 0
	for center := 0; center < len(s); center++ {
		ans += expand(s, center, center) + expand(s, center, center+1)
	}
	return ans
}

func expand(s string, left, right int) int {
	ans := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		ans++
		left--
		right++
	}
	return ans
}
