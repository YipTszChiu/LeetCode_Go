package binary_search

import "sort"

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	// 零值处理
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	flagL, indexL := searchLeftIndex_34(nums, target)
	if flagL == 0 {
		return []int{-1, -1}
	}

	if flagL == 1 { // 此时indexL为target在nums中的最左位置
		_, indexR := searchLeftIndex_34(nums, target+1)
		//if flagR == 0 {
		//	// 此时indexR为target+1应该插入的位置
		//	//  target 仅有一个
		//	if indexR-indexL == 1 {
		//		return []int{indexL, indexL}
		//	}
		//	// target 不仅有一个
		//	return []int{indexL, indexR - 1}
		//} else { // flagR == 1
		//	// 此时indexR为target+1在nums中的最左位置
		//	if indexR-indexL == 1 {
		//		//  target 仅有一个
		//		return []int{indexL, indexL}
		//	}
		//	// target 不仅有一个
		//	return []int{indexL, indexR - 1}
		//}

		// 上面发现 flagR 对于结果没有影响，分支的返回结果都是一致的，因此可以优化
		if indexR-indexL == 1 {
			return []int{indexL, indexL}
		}
		return []int{indexL, indexR - 1}
	}

	return []int{-1, -1}

}

// 两个功能：1.查找target是否存在；2.target插入nums[]的最左位置
func searchLeftIndex_34(nums []int, target int) (int, int) {
	left, right := 0, len(nums)-1

	// 表示nums[]中不存在target
	flag := 0

	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			flag = 1
			// 需要继续搜索左边有无target
			right = mid - 1
		} else if target < nums[mid] {
			right = mid - 1
		} else { // target > nums[mid]
			left = mid + 1
		}
	}

	return flag, left

}

// 调包
func searchRange_2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// sort.SearchInts() 查找 target 在数组中最左的位置，如果不存在，则返回可插入的位置
	left := sort.SearchInts(nums, target)
	// 注意 SearchInts函数可能返回len，因此也需要考虑等于len的情况
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums, target+1)
	return []int{left, right - 1}
}
