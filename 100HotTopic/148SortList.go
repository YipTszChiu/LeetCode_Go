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
