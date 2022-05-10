package _00HotTopic

import "sort"

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	// 提示 sort.SearchInts 函数如果碰到有相同值的元素，会返回第一个的位置
	leftIndex := sort.SearchInts(nums, target)
	// 判断 leftIndex 是否符合要求: 查找成功返回对应值的 Index ，但查找失败的话会返回 target 插入的位置
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}
	// 查找比 target 大 1 的位置
	rightIndex := sort.SearchInts(nums, target+1)
	return []int{leftIndex, rightIndex - 1}
}
