package _00HotTopic

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/?favorite=2cktkvj

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
