package _00HotTopic

// 438. 找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/?favorite=2cktkvj

// 滑动窗口
func findAnagrams(s string, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	// s长度小于p，s中不存在p的异位词
	if sLen < pLen {
		return
	}

	// 两个长度为26的数组存储字符数量
	var sCount, pCount [26]int

	// 对前 pLen 位字符进行数量统计
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	// 初始化的同时如果恰好字符一致，起始位 0 可以构成异位词
	if sCount == pCount {
		ans = append(ans, 0)
	}

	// index 从 0 遍历到 sLen-pLen（超过这个值s的长度无法构成异位词）
	for i, ch := range s[:sLen-pLen] {
		// 滑动窗口右移
		// 将最左端的字符退出
		sCount[ch-'a']--
		// 将最右端的下一个字符加入
		sCount[s[i+pLen]-'a']++
		// 如果字符数量一致，则构成异位词
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}

	return
}

// 滑动窗口的优化
func findAnagrams2(s string, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	// s长度小于p，s中不存在p的异位词
	if sLen < pLen {
		return
	}

	// 使用单个数组维护两个字符串中字符的数量
	count := [26]int{}
	// 初始化数组
	for i, ch := range p {
		// s 中有的字符则增加数量
		count[s[i]-'a']++
		// p 中有的字符则减少数量
		count[ch-'a']--
	}

	// differ 维护两个字符串不同的字符数
	differ := 0
	for _, c := range count {
		if c != 0 {
			differ++
		}
	}
	// 初始化结束后differ为0，恰好构成异位词
	if differ == 0 {
		ans = append(ans, 0)
	}

	// 从左到右遍历字符串s
	for i, ch := range s[:sLen-pLen] {
		// 滑动窗口右移
		if count[ch-'a'] == 1 {
			// ch 是最左端的字符此时离开窗口，如果count中这个字符为1，说明这个字符s和p中不同，则不同的字符减少一个
			differ--
		} else if count[ch-'a'] == 0 {
			// ch 是最左端的字符此时离开窗口，如果count中这个字符为0，说明这个字符s和p中相同，则不同的字符增加一个
			differ++
		}
		// 窗口左端字符退出
		count[ch-'a']--

		if count[s[i+pLen]-'a'] == -1 {
			// 即将进入窗口的字符count为-1，说明s中该字符与p中匹配，不同的字符减少一个
			differ--
		} else if count[s[i+pLen]-'a'] == 0 {
			// 即将进入窗口的字符count为-1，说明s中该字符与p中不匹配，不同的字符增加一个
			differ++
		}
		// 窗口右端进入字符
		count[s[i+pLen]-'a']++

		// 滑动窗口右移结束，判断是否s和p匹配
		if differ == 0 {
			ans = append(ans, i+1)
		}
	}

	return
}
