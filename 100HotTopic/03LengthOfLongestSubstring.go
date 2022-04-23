package _00HotTopic

// 3. 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// 该版本为自己写的，效率与剑指offer的48题有所差距，对比复习并思考差异

func lengthOfLongestSubstring(s string) int {
	// 初始化一个哈希表，存储 出现字符-上一次出现位置 的映射
	hashMap := make(map[byte]int)
	// res 存储出现过的最大长度, temp 存储当前遍历的长度
	res, temp := 0, 0
	// 遍历整个字符串
	for i := 0; i < len(s); i++ {
		// 判断当前字符是否出现过
		if lastIndex, ok := hashMap[s[i]]; !ok {
			// 如果没有出现过，存入 当前字符-index 的映射
			hashMap[s[i]] = i
			temp++
		} else {
			// 如果出现过，判断与上次出现位置的距离  abcdaba
			if i-lastIndex > temp {
				// 字符两次出现的距离大于当前遍历长度temp，说明未发生重复
				// 更新hashMap
				hashMap[s[i]] = i
				temp++
			} else {
				// 字符两次出现的距离小于等于当前遍历长度temp，说明发生重复
				// 更新hashMap
				hashMap[s[i]] = i
				// 更新temp
				temp = i - lastIndex
			}
		}
		// res 取最大的 temp
		if temp > res {
			res = temp
		}
	}

	return res
}
