package _00HotTopic

// 19. 删除链表的倒数第 N 个结点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1 2 3 4 5 n = 2 删除4
	// 初始化两个指针指向 head
	pre, cur := head, head
	// pre 指针提前向前移动 n 位
	for i := 0; i < n; i++ {
		pre = pre.Next
	}
	// 如果提前移动后 pre 直接指向 nil，说明需要删除的节点是头结点
	if pre == nil {
		return head.Next
	}
	// 正式遍历链表
	for pre.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}
	// 循环结束时 cur 指向需要删除元素的前一个节点，保存该节点
	pre = cur
	cur = cur.Next
	// pre.Next 直接指向 cur.Next，即需要删除元素的后一个节点
	pre.Next = cur.Next

	return head
}
