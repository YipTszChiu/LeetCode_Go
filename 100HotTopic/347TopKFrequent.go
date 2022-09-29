package _00HotTopic

import "container/heap"

// 347. 前 K 个高频元素
// https://leetcode.cn/problems/top-k-frequent-elements/?favorite=2cktkvj

// 桶
func topKFrequentBucket(nums []int, k int) []int {
	// 使用哈希表记录数字出现过的次数
	numsMap := make(map[int]int)
	// 遍历 nums 数组
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	// bucket[count][nums] 两个维度表示：出现频率为count次的数为nums
	// 有可能出现频率相同的数字，因此使用二位数组，即nums[]中可能不止一个数
	buckets := make([][]int, len(nums)+1)
	// 遍历哈希表
	for num, count := range numsMap {
		// 初始化第二维
		if len(buckets[count]) <= 0 {
			buckets[count] = make([]int, 0)
		}
		// 将 num 存入出现次数为 count 的 bucket
		buckets[count] = append(buckets[count], num)
	}
	// ret 用于返回结果
	ret := make([]int, 0)
	// buckets[]一维定义为出现的频次，因此从后往前遍历 buckets 可以获得最多的 k 个
	for i := len(buckets) - 1; i >= 0; i-- {
		// 将出现过的频次的数append到ret中，最后返回k个
		if len(buckets[i]) > 0 {
			ret = append(ret, buckets[i]...)
		}
	}
	return ret
}

// 堆
type Heap347 [][2]int

func (h Heap347) Len() int           { return len(h) }
func (h Heap347) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h Heap347) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap347) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *Heap347) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequentHeap(nums []int, k int) []int {
	// 使用哈希表记录数字出现过的次数
	numsMap := make(map[int]int)
	// 遍历 nums 数组
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	h := &Heap347{}
	heap.Init(h)
	for num, count := range numsMap {
		heap.Push(h, [2]int{num, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}
