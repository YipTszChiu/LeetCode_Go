package list

// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/

// 双指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 使用一个节点作为开头简化处理
	newHead := &ListNode{0, nil}
	// 用于遍历的三个指针
	ptr, ptr1, ptr2 := newHead, list1, list2
	// 对两个链表进行遍历
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			ptr.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			ptr.Next = ptr2
			ptr2 = ptr2.Next
		}
		ptr = ptr.Next
	}
	// 对剩余元素进行链接
	if ptr1 != nil {
		ptr.Next = ptr1
	}
	if ptr2 != nil {
		ptr.Next = ptr2
	}

	return newHead.Next
}
