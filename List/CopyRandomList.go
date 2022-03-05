package List


//Definition for a Node.
type Node struct {
	Val int
    Next *Node
	Random *Node
}


func copyRandomList(head *Node) *Node {
	//如果输入空链表，直接返回nil
	if head == nil {
		return nil
	}

	//在原来的链表（A→B→C）上构成一个复制的结构（A→A'→B→B'→ →C'）
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val,Next: node.Next}
	}

	//处理Random指针
	for node := head; node != nil; node = node.Next.Next {
		//注意node的Random有可能为空，此时node.Random.Next不存在
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	//将两个链断开
	//A→A'→B→B'→C→C'变为A→B→C以及A'→B'→C'
	headNew := head.Next
	for node := head; node != nil; node = node.Next{
		nodeNew := node.Next
		node.Next = nodeNew.Next
		if node.Next != nil {
			nodeNew.Next = node.Next.Next
		}
	}

	return headNew
}
