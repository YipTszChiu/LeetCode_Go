package _00HotTopic

// 128. 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/

// 哈希表
// 时间复杂度：O(n) n为数组的长度
// 空间复杂度：O(n) 哈希表需要存储n个数
func longestConsecutive(nums []int) int {
	// 思路：使用哈希表；
	// 优化：假设已有 x, x+1, x+2, …… x+n 的连续序列，此时突然不连续，
	// 如果从 x+1 开始枚举，得到 x+1 …… x+n 的序列一定小于 x 为起点的序列
	// 因此可以跳过这些情况，即判断 x 是否有前驱数 x-1， 只匹配没有前驱数的数字，就不会出现上述情况

	// 由于题目中 len(nums) >= 0，需要判空
	if len(nums) == 0 {
		return 0
	}
	// 使用 numSet 记忆出现过的数字
	numSet := map[int]bool{}
	// 遍历 nums 填充 numSet
	for _, num := range nums { // Go语言查漏：使用 range 遍历数组获得第一个参数是 i，第二个参数才是实际值
		numSet[num] = true
	}
	// longestStreak 记录最长序列，初始化为0
	longestStreak := 0
	// 遍历哈希表寻找最长序列
	for num := range numSet { // Go语言查漏：使用range 遍历 map 获得的是 key, value
		// 数 x 的前驱数 x-1 不存在
		if !numSet[num-1] {
			// 连续序列由 currentNum 开始，由于 num 没有前驱数，因此从 num 开始
			currentNum := num
			// 由于新的序列从 num 开始，使用 currentStreak 记录新的序列长度，初始为1
			currentStreak := 1
			// 从 currentNum 开始遍历，找到最长的一个序列
			for numSet[currentNum+1] {
				currentNum += 1
				currentStreak += 1
			}
			// 里循环结束时判断序列长度
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	// 外层循环需要 O(n) 的时间复杂度，只有当一个数是连续序列的第一个数的情况下才会进入内层循环，然后在内层循环中匹配连续序列中的数，
	// 因此数组中的每个数只会进入内层循环一次。根据上述分析可知，总时间复杂度为 O(n)

	return longestStreak
}
