package Stack_Queue

type MaxQueue struct {
	q []int
	p []int
}

//func Constructor() MaxQueue {
//	return MaxQueue{}
//}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.p[0]
}

func (this *MaxQueue) Push_back(value int)  {
	this.q = append(this.q, value)
	for len(this.p) > 0 && value > this.p[len(this.p) - 1] {
		this.p = this.p[: len(this.p) - 1]
	}
	this.p = append(this.p, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	if this.q[0] == this.p[0] {
		this.p = this.p[1: ]
	}
	value := this.q[0]
	this.q = this.q[1: ]
	return value
}

//作者：tuotuoli
//链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/solution/mian-shi-ti-59-ii-dui-lie-de-zui-da-zhi-dan-diao-z/

