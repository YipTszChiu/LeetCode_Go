package list

// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 在头节点前构建一个虚假节点
	dummy := &ListNode{Val: 0, Next: head}
	// 使用快慢指针实现一次遍历完成任务
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 此时 slow 指向要删除节点的前一个节点
	slow.Next = slow.Next.Next
	return dummy.Next
}
