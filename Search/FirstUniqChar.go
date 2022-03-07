package Search

//面试题50. 第一个只出现一次的字符
//https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

//哈希表的思想
func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}

	//由于s只包含小写字母，只需要一个长度26的int数组就足够
	charHash := make([]int, 26)
	//遍历字符串s
	for _, char := range s {
		charHash[char - 'a']++
	}
	for _, char := range s {
		if charHash[char - 'a'] == 1 {
			return byte(char)
		}
	}

	return ' '
}
