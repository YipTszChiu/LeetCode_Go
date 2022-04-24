package _00HotTopic

// 4. 寻找两个正序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

// 使用二分法在找到第k个最小的数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 计算两个数组的总长度
	totalLength := len(nums1) + len(nums2)
	// 判断奇偶
	if totalLength%2 == 1 {
		minIndex := totalLength / 2
		// 当总长度为奇数，比如长度为 5，中位数是 5/2 = 2，下标从0开始，实际上是第3个最小的数
		return float64(getKthElement(nums1, nums2, minIndex+1))
	} else {
		midIndex := totalLength / 2
		// 当总长度为偶数，比如长度为 6，中位数是 6/2 = 3，下标从0开始，实际上是第3、4个最小的数
		return float64(getKthElement(nums1, nums2, midIndex)+getKthElement(nums1, nums2, midIndex+1)) / 2.0
	}
	return 0
}

// 从两个数组中找到第K个最小的数
func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		// 特例处理：1.越界；2.K为1
		// 当 index1 越界， 直接从 nums2 数组返回第K小的数
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		// index2 越界同理
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		// k为1，直接返回数组第一个元素
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		// 将 k 分为两半，分别到两个nums数组比较大小，取 k/2 - 1位置的数（推理过程见LeetCode官方题解）
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		//为什么越界仍然要取数组最后一个元素？因为有可能出现1，2，99这种情况，一个很大的元素排在了数组末尾
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			// +1的原因是包括当前newIndex位在内可以全部排除
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
