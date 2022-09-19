package _00HotTopic

// 287. 寻找重复数
// https://leetcode.cn/problems/find-the-duplicate-number/

// 快慢指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func findDuplicate(nums []int) int {
	// 将数组构建成 i → nums[i] 的图，由于 1 <= nums[i] <= n，且有一个重复数，因此该图必有环
	// 使用快慢指针算法查找环的起点（Floyd 判圈算法），141 与 142题
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	// 结束循环时 slow == fast，将一个指针移到头部按step为1进行遍历，交点即为环起点，在该题中为重复数
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
