package _00HotTopic

import "sort"

// 169. 多数元素
// https://leetcode.cn/problems/majority-element/?favorite=2cktkvj

// 使用哈希表
// 时间复杂度：O(n) 取决于数组元素
// 空间复杂度：O(n) 取决于数组元素
func majorityElement(nums []int) int {
	numSet := map[int]int{}
	length := len(nums)
	// 遍历 nums
	for _, elem := range nums {
		numSet[elem]++
		if numSet[elem] > length/2 {
			return elem
		}
	}

	return 0
}

// 排序法
// 时间复杂度：O(nlogn) 取决于排序算法
// 空间复杂度：O(1)
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Boyer-Moore 投票算法
// 时间复杂度：O(nlogn) 取决于数组长度
// 空间复杂度：O(1)
func majorityElement3(nums []int) int {
	candidate := -1
	count := 0
	for _, elem := range nums {
		if elem == candidate {
			count++
		} else if count-1 < 0 {
			candidate = elem
			count = 1
		} else {
			count--
		}
	}
	return candidate
}
