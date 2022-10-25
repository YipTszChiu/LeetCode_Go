package list

// 86. 分隔链表
// https://leetcode.cn/problems/partition-list/

// 双指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func partition(head *ListNode, x int) *ListNode {
	// 将原链表分隔成两个
	newHead1 := &ListNode{Val: 0, Next: nil}
	newhead2 := &ListNode{Val: 0, Next: nil}
	// 使用指针分别对大于等于x，小于x的值进行处理
	ptr, ptr1, ptr2 := head, newHead1, newhead2

	// 遍历原链表，将小于x的节点接入新链1，大于等于x的节点接入新链2
	for ptr != nil {
		if ptr.Val < x {
			ptr1.Next = ptr
			ptr1 = ptr1.Next
		} else {
			ptr2.Next = ptr
			ptr2 = ptr2.Next
		}
		ptr = ptr.Next
	}
	// 遍历结束后将新链1的尾部和新链2的头部相接，并断开新链2的尾部
	ptr1.Next = newhead2.Next
	ptr2.Next = nil
	// 返回分隔后的头节点
	return newHead1.Next
}

