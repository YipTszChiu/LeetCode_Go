package Stack_Queue

import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}


//func Constructor() CQueue {
//	return CQueue{
//		stack1:list.New(),
//		stack2:list.New(),
//	}
//}


func (this *CQueue) AppendTail(value int)  {
	//直接入栈1
	this.stack1.PushBack(value)
}


func (this *CQueue) DeleteHead() int {
	//第二个栈为空时，将第一个栈的元素出栈到第二个栈
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}

	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}

	//stack2仍然为空
	return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
