package array

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/

// 双指针法
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expand(s, i, i)
		left2, right2 := expand(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
