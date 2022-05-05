package _00HotTopic

// 31. 下一个排列
// https://leetcode-cn.com/problems/next-permutation/

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 从后往前遍历，找到第一个非降序的数字
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// i 如果 < 0，说明整个数组已经是降序的，因此已经是最大排列
	// 遍历后找到第一个非降序数字nums[i]
	if i >= 0 {
		j := n - 1
		// 继续从后往前遍历，找到第一个比nums[i]大的数
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		// 交换找到的两个数的位置
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 对 i+1 之后的位置进行逆序
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
