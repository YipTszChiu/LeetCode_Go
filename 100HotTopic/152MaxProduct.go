package _00HotTopic

// 152. 乘积最大子数组
// https://leetcode.cn/problems/maximum-product-subarray/

// 动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProduct(nums []int) int {
	/*	思路：类似最大自组和，使用 dp 维护子组最大乘积
		但存在一个问题：如果前面子组乘积为负数，后面碰到负数会有更大的乘积
		因此除了维护最大乘积，还要维护最小乘积
	*/
	length := len(nums)
	// 初始化最大乘积和最小乘积两个数组
	maxMul, minMul := make([]int, length), make([]int, length)
	maxMul[0], minMul[0] = nums[0], nums[0]
	res := maxMul[0]
	// 遍历 nums 计算结果
	for i := 1; i < length; i++ {
		// 最大乘积：取 1.前一位最大乘积与nums[本位]相乘，2.nums[本位]，以及 3.前一位最小乘积与nums[本位]相乘 三者的最大值
		//						正数 * 正数									负数 * 负数
		maxMul[i] = max(maxMul[i-1]*nums[i], max(nums[i], minMul[i-1]*nums[i]))
		// 最小乘积：取 1.前一位最小乘积与nums[本位]相乘，2.nums[本位]，以及 3.前一位最大乘积与nums[本位]相乘 三者的最小值
		//						负数 * 正数									正数 * 负数
		minMul[i] = min(minMul[i-1]*nums[i], min(nums[i], maxMul[i-1]*nums[i]))
		if res < maxMul[i] {
			res = maxMul[i]
		}
	}

	return res
}

// 空间优化
func maxProduct2(nums []int) int {
	length := len(nums)
	// 初始化最大乘积和最小乘积两个数组
	maxMul, minMul, res := nums[0], nums[0], nums[0]
	// 遍历 nums 计算结果
	for i := 1; i < length; i++ {
		// 需要存储一下maxMul，避免执行过程中maxMul改变后，minMul计算受到影响
		mA := maxMul
		maxMul = max(maxMul*nums[i], max(nums[i], minMul*nums[i]))
		minMul = min(minMul*nums[i], min(nums[i], mA*nums[i]))
		if res < maxMul {
			res = maxMul
		}
	}
	return res
}
