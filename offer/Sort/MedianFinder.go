package Sort

import "container/heap"

//剑指 Offer 41. 数据流中的中位数
//https://xd2r8gb0a1.feishu.cn/docs/doccnV7SizUChKHwxjtbKWJ35Zb#

//heap需要实现Len、Less、Swap；Push、Pop
type maxHeapp []int //同包内有另一个maxHeap，因此多打一个p以区分
type minHeapp []int

func (m *maxHeapp) Len() int {
	return len(*m)
}

func (m *minHeapp) Len() int {
	return len(*m)
}

//大根堆需要arr[i] > arr[j]
func (m *maxHeapp) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

//小根堆需要arr[i] < arr[j]
func (m *minHeapp) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *maxHeapp) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapp) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeapp) Push(x interface{}) {
	(*m) = append((*m), x.(int))
}

func (m *minHeapp) Push(x interface{}) {
	(*m) = append((*m), x.(int))
}

func (m *maxHeapp) Pop() (x interface{}) {
	*m, x = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

func (m *minHeapp) Pop() (x interface{}) {
	*m, x = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

//使用一个大根堆和一个小根堆维护
type MedianFinder struct {
	maxH *maxHeapp //大根堆用于存储较小一半的数据
	minH *minHeapp //小根堆用于存储较大一半的数据
	//注意有可能会有奇数个数据，规定当数据个数为奇数时，minHeap.Len() > maxHeap.Len()，即多出的一个插入到minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		new(maxHeapp),
		new(minHeapp),
	}
}

func (this *MedianFinder) AddNum(num int) {
	//当大小根堆长度相等，即插入的num造成奇数个数据，根据规定插入到小根堆
	if this.maxH.Len() == this.minH.Len() {
		//如果是空堆，直接插入
		if this.minH.Len() == 0 || num >= (*this.minH)[0] {
			heap.Push(this.minH, num)
		} else {
			//如果不是空堆，需要判断:
			//1.num >= min堆顶，直接插入到min堆（合并到上一步）
			//2.num < min堆顶，即num实际应该插入max堆

			//插入到max堆，此时max堆比min堆多一个元素，需要调整至符合规定
			heap.Push(this.maxH, num)
			//max堆弹出堆顶，将堆顶插入到min堆，此时符合奇数个元素时，min堆比max堆多一
			top := heap.Pop(this.maxH).(int)
			heap.Push(this.minH, top)
		}
	} else {
		//大小根堆长度不等，即此时符合max堆比min堆少一，直接插入到max堆，但仍需判断：
		//1.num > min堆顶，将num插入到min堆
		if num > (*this.minH)[0] {
			heap.Push(this.minH, num)
			//此时min堆比max堆多2，从min堆将堆顶弹出，插入到max堆
			top := heap.Pop(this.minH)
			heap.Push(this.maxH, top)
		} else {
			//2.num <= max堆顶，直接插入到max堆
			heap.Push(this.maxH, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	//偶数个元素，返回两个中间数和的二分一
	if this.maxH.Len() == this.minH.Len() {
		return float64((*this.maxH)[0])/float64(2) + float64((*this.minH)[0])/float64(2)
	} else {
		//奇数个元素，返回中位数
		return float64((*this.minH)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
