package _00HotTopic

import "sort"

// 15. 三数之和
// https://leetcode-cn.com/problems/3sum/

// 排序+双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	// 对数组进行排序
	sort.Ints(nums) // O(NlogN)
	res := make([][]int, 0)
	// 三重循环遍历：注意使用两个指针往中间夹逼的方法优化
	for first := 0; first < n; first++ {
		// 题目要求不能有重复的序列，碰到相同的数字则跳过
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		// 同样，第二位如果重复也跳过
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 第三位从数组末端开始左移
			for third > second && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
