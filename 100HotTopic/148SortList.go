package _00HotTopic

// 148. 排序链表
// https://leetcode.cn/problems/sort-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 自顶向下归并排序
// mergeList 用于合并两个有序链表
func mergeList(head1, head2 *ListNode) *ListNode {
	// 创建一个新的空链
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	// 遍历两个有序链表
	for temp1 != nil && temp2 != nil {
		// 元素更小的接到新的链
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	// 检查剩余元素并加入到新的链后
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

// 通过递归法，从中间点分成两个链表归并排序
func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	// slow fast分别以步长为 1 2向后移动
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	// 遍历结束时 slow即为链表中点
	mid := slow
	// 归并排序：注意这里 mid
	return mergeList(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

// 自底向上归并排序
func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 记录链表的长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	// 自底向上：子链从 1 开始，每次翻倍
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		// 按照 subLength 的长度从头到位进行归并
		for cur != nil {
			// 选取subLength个元素为第一个链
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 第二个链头为上面遍历结束的下一个节点
			head2 := cur.Next
			// 断开两个子链
			cur.Next = nil
			// 选取subLength个元素为第二个链
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 断开第二个子链以及后面未划分的链
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			// 合并两个有序链表
			prev.Next = mergeList(head1, head2)
			// 合并结束，将 prev 指针指向已经归并完的链的末尾
			for prev.Next != nil {
				prev = prev.Next
			}
			// cur 指向未归并的链准备下一次归并
			cur = next
		}
	}
	return dummyHead.Next
}
