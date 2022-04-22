package _00HotTopic

// 2. 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 由于链表是数字的逆序，因此head指向两个数字的个位
	// 初始化一个head，作为计算结果的头指针
	head := &ListNode{
		0,
		nil,
	}
	// 使用变量 c 存储每一位相加的进位
	c := 0
	// 初始化一个 current 指针，指向计算的结果位，初始指向头指针
	current := head

	// 对 l1 和 l2 进行遍历相加
	for l1 != nil && l2 != nil {
		// temp 暂存相加后的结果（个位）
		temp := (l1.Val + l2.Val + c) % 10
		// c为进位
		c = (l1.Val + l2.Val + c) / 10
		// 新建一个节点接到 current 之后
		current.Next = &ListNode{
			temp,
			nil,
		}
		// current, l1, l2 后移
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	// 循环结束后判断 l1, l2 是否遍历完成
	if l1 != nil {
		for l1 != nil {
			temp := (l1.Val + c) % 10
			c = (l1.Val + c) / 10
			current.Next = &ListNode{
				temp,
				nil,
			}
			current = current.Next
			l1 = l1.Next
		}
	}
	if l2 != nil {
		for l2 != nil {
			temp := (l2.Val + c) % 10
			c = (l2.Val + c) / 10
			current.Next = &ListNode{
				temp,
				nil,
			}
			current = current.Next
			l2 = l2.Next
		}
	}

	// 判断是否有多余进位
	if c != 0 {
		current.Next = &ListNode{
			c,
			nil,
		}
		// 由于计算结束，current可以不后移
	}

	return head.Next
}
