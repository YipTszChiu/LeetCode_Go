package mySort

func MergeSort(arr []int) []int {
	length := len(arr)
	// 当数组分割至只剩一个元素，直接返回
	if length < 2 {
		return arr
	}
	// 对数组进行分割
	mid := length / 2
	left := arr[:mid]
	right := arr[mid:]
	// 每次分割成两半，并对这两半进行排序，然后返回排序后的数组
	return mergeSort(MergeSort(left), MergeSort(right))

}

func mergeSort(left, right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		res = append(res, right[0])
		right = right[1:]
	}

	return res
}
