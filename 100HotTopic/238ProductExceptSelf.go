package _00HotTopic

// 238. 除自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/?favorite=2cktkvj

// 上下三角
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}

	// 循环左下三角
	for i := 1; i < len(result); i++ {
		result[i] = nums[i-1] * result[i-1]
	}
	// 循环右上三角
	mul := 1
	for i := len(result) - 2; i >= 0; i-- {
		mul *= nums[i+1]
		result[i] *= mul
	}

	return result
}
