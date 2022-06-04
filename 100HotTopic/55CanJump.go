package _00HotTopic

// 55.跳跃游戏
// https://leetcode.cn/problems/jump-game/

func canJump(nums []int) bool {
	n := len(nums)
	// 维护能够到达的最远端，初始在 nums[0]
	rightMost := 0
	// 对每个元素进行遍历
	for i := 0; i < n; i++ {
		if i <= rightMost {
			// 更新能够达到的最远端，取最大值
			rightMost = max(rightMost, i+nums[i])
			// 如果 rightMost >= n-1 说明可以达到最后一个下标
			if rightMost >= n-1 {
				return true
			}
		}
	}
	return false
}
