package _00HotTopic

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/

// 枚举法
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func subarraySum(nums []int, k int) int {
	// count 记录返回值
	count := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// 前缀和+哈希表
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func subarraySum2(nums []int, k int) int {
	// pre[i] = nums[0] + nums[1] + ... + nums[i]
	// pre[i] = pre[i-1] + nums[i]
	// 题目所要求的子组 j~i 的和为 k，即 j + (j+1) + ... + i == k ?
	// 可以转化为 pre[i] - pre[j-1] == k ?
	// 移项后 pre[i] - k == pre[j-1]，即遍历到 i 时，之前是否出现过和为 pre[i]-k 的 pre[j]，如果有则和为k的子组加一

	count, pre := 0, 0
	// 使用哈希表记录前缀和
	hashMap := make(map[int]int)
	// 初始状态和为 0，出现一次；为什么需要和为0？第一次遍历到pre[i] = k，pre[i]-k此时就等于0，因此需要记录一次
	hashMap[0] = 1
	// 从左到右更新 hashMap 就能保证 pre[j] 出现在 pre[i] 之前，可以满足pre[i] - k == pre[j]
	for i := 0; i < len(nums); i++ {
		// pre 即为上面所说的 pre[i]
		pre += nums[i]
		// 检查是否有 pre[i]-k == pre[j] 出现过
		if _, ok := hashMap[pre-k]; ok {
			// 如果出现过，则子组和等于k的次数需要加其出现次数
			count += hashMap[pre-k]
		}
		hashMap[pre]++
	}
	return count
}
