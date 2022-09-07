package _00HotTopic

// 215. 数组中的第K个最大元素
// https://leetcode.cn/problems/kth-largest-element-in-an-array/?favorite=2cktkvj

// 方法一：快速排序优化算法
// 思路： 快排可以确定一个数的最终index，并且这个index左边比该数小，右边比该数大，因此对比index和k，只需要考虑半边
func findKthLargest(nums []int, k int) int {
	topKSplit(nums, k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

// 快速排序并返回本次排序确定的index
func quickSort(nums []int, start, stop int) int {
	if start >= stop {
		return -1
	}
	// 定下pivot用于排序，并存储该值
	pivot := nums[start]
	l, r := start, stop
	for l < r {
		// 从右边开始找到第一个小于 pivot 的数
		for l < r && nums[r] >= pivot {
			r--
		}
		// 将这个小于 pivot 的数放到左边
		nums[l] = nums[r]
		// 从左边开始找到第一个大于 pivot 的数
		for l < r && nums[l] < pivot {
			l++
		}
		// 将这个大于 pivot 的数放到右边
		nums[r] = nums[l]
	}
	// 循环结束 l == r，这个位置即为 pivot 的 index
	nums[l] = pivot
	return l
}

// topKSplit 调用 quickSort 排序并找到第k大的数
func topKSplit(nums []int, k, start, stop int) {
	if start < stop {
		// 进行一次快排并获取此次排序已经确定好的index
		index := quickSort(nums, start, stop)
		// 注意排序是升序，index 为 len(nums)-k 即第k大数
		if index == len(nums)-k {
			return
		} else if index < len(nums)-k {
			// index 在目标第k大数左边，对右边的切片排序
			topKSplit(nums, k, index+1, stop)
		} else {
			// index 在目标第k大数右边，对左边的切片排序
			topKSplit(nums, k, start, index-1)
		}
	}
}

// 方法二：小根堆
// 思路：构建一个容量为 k 的小根堆，堆顶即为k个元素里最小的，然后遍历数组将所有元素进堆。
// 由于容量为 k 最小的数在遍历时已经被 pop，剩余的k个数是最大的k个，而堆顶则是k个里最小的，即第k大数
type Heap struct {
	arr  []int
	size int
}

func (hp *Heap) Add(num int) {
	// 没有超过堆的容量，直接 append
	if len(hp.arr) < hp.size {
		hp.arr = append(hp.arr, num)
		// append 后对新增节点进行排序
		for i := len(hp.arr) - 1; i > 0; {
			// i 的父节点
			p := (i - 1) / 2
			// 父节点大于子节点，需要交换
			if p >= 0 && hp.arr[p] > hp.arr[i] {
				hp.arr[p], hp.arr[i] = hp.arr[i], hp.arr[p]
				// 继续往上遍历
				i = p
			} else {
				break
			}
		}
	} else if num > hp.arr[0] {
		// 超过堆容量，将堆顶进行替换，然后重新排序
		hp.arr[0] = num
		hp.Down(0)
	}
}

func (hp *Heap) Down(i int) {
	k := i
	// l, r 为 i 的左右子节点
	l, r := 2*i+1, 2*i+2
	n := len(hp.arr)
	// 比较父节点与两个叶子节点，选最小的一个
	if l < n && hp.arr[k] > hp.arr[l] {
		k = l
	}
	if r < n && hp.arr[k] > hp.arr[r] {
		k = r
	}
	// 如果需要交换，则父节点与最小的一个交换
	if i != k {
		hp.arr[i], hp.arr[k] = hp.arr[k], hp.arr[i]
		hp.Down(k)
	}
}

func findKthLargest1(nums []int, k int) int {
	hp := &Heap{size: k}
	for _, num := range nums {
		hp.Add(num)
	}
	return hp.arr[0]
}
