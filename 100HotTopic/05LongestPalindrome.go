package _00HotTopic

// 5. 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/

// 解法一：动态规划
func longestPalindrome(s string) string {
	length := len(s)
	// 当字符串只有一个字符的时候，必定是回文串
	if length < 2 {
		return s
	}

	maxLen, begin := 1, 0

	// dp[i][j] 表示 s[i:j] 是否为回文串
	dp := [][]bool{}
	// 完成dp的构造，并且将dp[i][i]初始化
	for i := 0; i < length; i++ {
		temp := make([]bool, length)
		temp[i] = true
		dp = append(dp, temp)
	}

	// 枚举子串长度，1已经在开头直接返回，因此从2开始
	for l := 2; l <= length; l++ {
		// 枚举左边界
		for i := 0; i < length; i++ {
			// 由子串长度 l 和左边界 i 可以得出右边界：j - i + 1 = l
			j := l + i - 1
			// 右边界越界则退出当前循环
			if j >= length {
				break
			}
			// 判断左右边界的两个字符是否一致
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				// 如果一致，且子串长度不超过（<=）3个字符，则dp[i][j]是回文串，置为true
				if j-i < 3 {
					dp[i][j] = true
				} else {
					// 如果一致，但子串长度超过3个字符，则dp[i][j]是否为回文串取决于dp[i+1][j-1]
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 如果当前子串是回文串，且长度比记录的回文串大，则更新信息
			if dp[i][j] == true && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

// 解法二：中心拓展法
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	// 左边界从0开始遍历
	for i := 0; i < len(s); i++ {
		// 参数left right均为i，用于判断奇数长度的回文
		left1, right1 := expandAroundCenter(s, i, i)
		// 参数left right相差1，用于判断偶数长度的回文
		left2, right2 := expandAroundCenter(s, i, i+1)
		// 取较长的回文作为结果
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	// 拓展条件：1.left right不越界；2.且s[left] == s[right]
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
