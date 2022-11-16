package mySort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		// 以第一个数作为基准进行分治
		pivot := arr[left]
		i, j := left, right

		for i < j {
			// 从右往左找到第一个小于等于 pivot 的数
			for i < j && arr[j] > pivot {
				j--
			}
			// 将这个数与左索引互换
			if i < j {
				arr[i] = arr[j]
				i++
			}
			// 从左往右找到第一个大于等于 pivot 的数
			for i < j && arr[i] < pivot {
				i++
			}
			// 将这个数与右索引互换
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}

		// 循环结束后指针相遇在 pivot 实际位置
		arr[i] = pivot
		// 递归
		quickSort(arr, left, i-1)
		quickSort(arr, i+1, right)
	}
}
