package _00HotTopic

// 23. 合并K个升序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法一：顺序合并——效率较低
func mergeKLists(lists []*ListNode) *ListNode {
	// 使用 res 变量维护一个链表，每次把一个 lists 中的链表与 ans 合并
	var res *ListNode
	res = nil
	k := len(lists)
	for i := 0; i < k; i++ {
		res = mergeTwoLists2(res, lists[i])
	}
	return res
}

func mergeTwoLists2(a, b *ListNode) *ListNode {
	// 传入空链表，直接返回另一个链表
	if a == nil || b == nil {
		if a == nil {
			return b
		} else {
			return a
		}
	}
	// 初始化一个头结点以及尾节点
	head := &ListNode{0, nil}
	tail := head
	// 合并两个有序链表
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	// 结束后检查两个链表是否遍历完全
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return head.Next
}

// 方法二：分治合并——类似归并排序的思路
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists2(merge(lists, l, mid), merge(lists, mid+1, r))
}
