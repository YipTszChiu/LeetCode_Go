package _00HotTopic

// 53. 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	// 初始化 maxSum 为 nums[0]
	maxSum := nums[0]
	// 通过 dp规划 的方式进行遍历
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}

	return maxSum
}
