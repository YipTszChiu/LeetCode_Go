package DivideAndConquer

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

//递归实现归并排序
func mergeSort(arr []int, start, end int) int {
	if start >= end {
		return 0
	}

	// 将数组一分为二
	mid := start + (end-start)/2
	//分别排序左右两个子数组
	count := mergeSort(arr, start, mid) + mergeSort(arr, mid+1, end)

	//temp用于暂存归并排序的子数组
	temp := []int{}
	//左右指针初始化指向start以及mid+1
	left, right := start, mid+1

	// 循环条件为左右指针都不越界
	for left <= mid && right <= end {
		// 左指针元素小于右指针元素
		if arr[left] <= arr[right] {
			//将小的加入到暂存数组
			temp = append(temp, arr[left])
			//逆序贡献记录到count
			count += right - (mid + 1)
			left++
		} else {
			// 右指针元素小于左指针元素
			//将小的加入到暂存数组
			temp = append(temp, arr[right])
			right++
		}
	}

	// 循环结束后检查左右是否遍历完全
	for ; left <= mid; left++ {
		temp = append(temp, arr[left])
		count += end - (mid + 1) + 1
	}
	for ; right <= end; right++ {
		temp = append(temp, arr[right])
	}

	// 用暂存数组替换arr
	for i := start; i <= end; i++ {
		arr[i] = temp[i-start]
	}

	return count
}
