package DynamicProgramming

//剑指 Offer 48. 最长不含重复字符的子字符串
//https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

//哈希表+dp
func lengthOfLongestSubstring(s string) int {
	//构建哈希表，并初始化：ASCII码有128个
	charMap := make([]int, 128)
	for i := 0; i < len(charMap); i++ {
		//charMap的值表示一个字符上次出现的位置，初始化值为-1
		charMap[i] = -1
	}

	//max为最大距离，temp为临时距离，lastIndex为一个字符上次出现的位置
	max, temp, lastIndex := 0, 0, -1
	for i := 0; i < len(s); i++ {
		//获取s[i]字符上次出现的位置
		lastIndex = charMap[s[i]]
		//将s[i]字符当前位置记录到charMap
		charMap[s[i]] = i
		//字符串重复的距离比临时距离大，说明temp内没有发生重复
		if temp < i-lastIndex {
			temp++
		} else {
			//temp >= i-lastIndex
			//证明在temp临时距离内，出现了相同字符，因此最长不含重复距离为i-lastIndex
			temp = i - lastIndex
		}
		//更新max
		max = Max(temp, max)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}