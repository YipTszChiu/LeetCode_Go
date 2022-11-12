package Sort

import "container/heap"

//剑指 Offer 40. 最小的k个数
//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

//法一：直接排序法；效率较低
//func getLeastNumbers(arr []int, k int) []int {
//	//空值处理
//	if len(arr) == 0 || arr == nil || k == 0 {
//		return nil
//	}
//	//对arr进行排序
//	mySort.Ints(arr)
//	//res用于返回结果
//	res := []int{}
//	//遍历排序后的arr将最小的k个数加入到res
//	for i := 0; i < k; i++ {
//		res = append(res, arr[i])
//	}
//
//	return res
//}

// 法二：大根堆维护k个数
// 使用heap需要实现Len(),Less(),Swap(),Push(),Pop()
type maxHeap []int

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() (x interface{}) {
	*m, x = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

// Peek用于查看堆顶
func (m *maxHeap) Peek() interface{} {
	return (*m)[0]
}

func getLeastNumbers(arr []int, k int) []int {
	//res用于返回结果
	res := []int{}
	//空值处理
	if len(arr) == 0 || arr == nil || k == 0 {
		return res
	}

	//构建一个大根堆
	mh := &maxHeap{}
	//在大根堆mh中调用Push插入k个值
	for i := 0; i < k; i++ {
		heap.Push(mh, arr[i])
	}
	//初始化大根堆
	heap.Init(mh)

	n := len(arr)
	//对剩下的元素进行遍历
	for i := k; i < n; i++ {
		//如果元素小于大根堆的堆顶，则需要将当前堆顶弹出，将当前元素插入大根堆中
		if mh.Len() > 0 && arr[i] < mh.Peek().(int) {
			heap.Pop(mh)
			heap.Push(mh, arr[i])
		}
	}

	//遍历结束后将大根堆中的元素加入到res中
	for i := 0; i < k; i++ {
		//注意这样Pop得到的res是从大到小排列
		res = append(res, heap.Pop(mh).(int))
	}
	return res
}
