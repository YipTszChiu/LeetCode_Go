package _00HotTopic

func canPartition(nums []int) bool {
	n := len(nums)
	// 数组元素少于2无法分割成两半
	if n < 2 {
		return false
	}
	// max 记录数组最大元素
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	// 数组和为奇数无法分割成两半
	if sum%2 != 0 {
		return false
	}
	// 目标是找到元素和为数组和的一半，则可以分割
	target := sum / 2
	// 数组最大元素大于和的一半，无法分割
	if max > target {
		return false
	}
	// dp[i][j] 表示在 nums[0]-nums[i] 中选取任意个，可以构成和为j
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 不选取任何数则可以构成和为0，因此dp[i][0]全为真
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	// 选取 nums[0] 可以构成和为 nums[0]
	dp[0][nums[0]] = true

	for i := 1; i < n; i++ {
		// 对于当前遍历的元素 nums[i]
		v := nums[i]
		for j := 1; j <= target; j++ {
			// 如果和j比当前元素大
			if j >= v {
				// 可以不选择当前元素作为划分：dp[i][j] = dp[i-1][j]，即仍在0-(i-1)个元素中任意选取，组成和为j
				// 也可以选择当前元素作为划分：dp[i][j] = dp[i-1][j-v]，即在0-(i-1)个元素中选取和为j-v，再选取nums[i] 组成和为 j
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				// 如果和j比当前元素小，无法选取nums[i]，dp[i][j] = dp[i-1][j]
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}
