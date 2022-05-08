package _00HotTopic

// 33. 搜索旋转排序数组
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

// 自己的做法，效率低且可读性差，看下面的官方解法
func search(nums []int, target int) int {
	// 保存 nums 的长度
	n := len(nums)
	// 对 n = 0 或 n = 1 的特例进行处理
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 初始化左右指针 left 和 right 对 nums 进行二分查找
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		// 如果二分查找的 mid 位置就是 target，直接返回下标 mid
		if nums[mid] == target {
			return mid
		}
		// 判断从 0 到 mid 的单调性，单调递增则进入该if
		if nums[0] <= nums[mid] {
			// 判断 mid 与 target 的关系
			if nums[mid] < target {
				// target 大于 mid，又由于 0-mid 单调递增，因此 target只可能出现在右半部分
				left = mid + 1
			} else if target < nums[0] {
				// target 小于 mid，且 target 小于 nums[0]，说明target只可能出现在右半部分
				left = mid + 1
			} else {
				// target 小于 mid，但target 大于等于 nums[0]，说明target只可能出现在左半部分
				right = mid - 1
			}
		} else {
			// 0 - mid 非单调递增，或者说 mid - (n-1)单调递增
			// 判断 mid 与 target 的关系
			if nums[mid] < target {
				// target 大于 mid
				if target <= nums[n-1] {
					// 且 target 不大于 nums[n-1]，则 target 只可能出现在右半部分
					left = mid + 1
				} else {
					// target 大于 mid 但 target 大于 nums[n-1]，则 target 只可能出现在左半部分
					right = mid - 1
				}
			} else {
				// target 小于 mid，target 只可能出现在左半部分
				right = mid - 1
			}
		}
	}

	return -1
}

func search2(nums []int, target int) int {
	// 保存 nums 的长度
	n := len(nums)
	// 对 n = 0 或 n = 1 的特例进行处理
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 初始化左右指针 left 和 right 对 nums 进行二分查找
	left, right := 0, n-1

	for left <= right {
		// 计算中间点
		mid := (left + right) / 2
		// nums[mid] 即为 target 则直接返回
		if nums[mid] == target {
			return mid
		}
		// 0-mid 单调增   注意这个等号，试想输入[3,1] target = 1
		if nums[0] <= nums[mid] {
			// target 比单调增区间的左端点大，且 target 比单调区间右端点小，说明 target 落在 0-mid 中
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // mid-(n-1) 单调增
			// target 比单调增区间的右端点小，且 target 比单调增区间的左端点大，说明 target 落在 mid - (n-1) 中
			if nums[n-1] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
