package _00HotTopic

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/?favorite=2cktkvj

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 迭代法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 使用两个指针存储
	var pre *ListNode
	curr := head
	for curr != nil {
		// 需要记录 next 指针避免断链
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

// 递归法
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 先一直走到链尾
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
