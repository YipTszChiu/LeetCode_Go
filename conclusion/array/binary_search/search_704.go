package binary_search

// 704. 二分查找
// https://leetcode.cn/problems/binary-search/

func search(nums []int, target int) int {
	// 0 1 2 3
	// 0 1 2 3 4

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
