package array

// 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	left, right := 0, 1
	for right < len(nums) {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
		right++
	}

	return left + 1
}
