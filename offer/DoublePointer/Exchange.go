package DoublePointer

//剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
//https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		//从前往后找到第一个偶数
		for nums[i]%2 == 1 && i < j {
			i++
		}
		//从后往前找到第一个奇数
		for nums[j]%2 == 0 && i < j {
			j--
		}
		//奇偶交换
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
