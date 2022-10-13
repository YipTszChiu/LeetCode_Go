package _00HotTopic

// 581. 最短无序连续子数组
// https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/?favorite=2cktkvj

// 双指针 + 线性扫描
// 时间复杂度：O(n) 只需要一次遍历
// 空间复杂度：O(1)
func findUnsortedSubarray(nums []int) int {
	// 原数组可以看作：左边以及右边的部分单调递增，中间一部分乱序
	// 但中间乱序部分有可能存在比左边右端点还小的数，或比右边左端点还大的数，只是调整中间乱序部分的顺序并不够，需要找到这种点并调整
	n := len(nums)
	l, r := 0, n-1
	// 首先找到左右两个有序端点
	for ; l < r && nums[l] <= nums[l+1]; l++ {
	}
	for ; l < r && nums[r] >= nums[r-1]; r-- {
	}
	// 记录乱序的起止 index
	start, end := l, r
	// 记录有序端点值
	min, max := nums[l], nums[r]
	// 在乱序部分查找是否有小于min或大于max的点
	for i := start; i <= end; i++ {
		if nums[i] < min {
			// 乱序部分存在 nums[i] 比左端点小，需要移动左端点
			for ; l >= 0 && nums[i] < nums[l]; l-- {
			}
			min = nums[i]
		}
		if nums[i] > max {
			// 乱序部分存在 nums[i] 比右端点大，需要移动右端点
			for ; r < n && nums[i] > nums[r]; r++ {
			}
			max = nums[i]
		}
	}

	if l == r {
		return 0
	}
	return (r - 1) - (l + 1) + 1
}
