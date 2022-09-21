package _00HotTopic

// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/?favorite=2cktkvj

// 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	// dp[i] 表示以第i个数字结尾的最长上升子序列的长度，nums[i] 必须被选取
	// 转移方程 dp[i] = max(dp[j]) + 1，0 <= j < i 且 nums[j] < nums[i]
	dp := make([]int, len(nums))
	dp[0] = 1
	maxCount := 1
	for i := 1; i < len(nums); i++ {
		// 将默认值改为 1，因为最短子序列是其本身
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxCount = max(maxCount, dp[i])
	}

	return maxCount
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
