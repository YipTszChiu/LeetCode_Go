package _00HotTopic

import "math"

// 155.最小栈
// https://leetcode.cn/problems/min-stack/

type MinStack struct {
	// 使用一个切片实现最基础的栈
	stack []int
	// 另一个切片辅助实现最小栈的功能
	minStack []int
}

// 构造器
func Constructor_MinStack() MinStack {
	return MinStack{
		stack: []int{},
		// 初始化时 minStack 栈底存一个最大的整数，便于函数处理
		minStack: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	minTop := len(this.minStack)
	// val 入栈前与栈内最小元素比较，如果更小则替换,否则仍然为上一位最小的元素
	if val < this.minStack[minTop-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[minTop-1])
	}
	// val 入栈
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	// 正常出栈
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
