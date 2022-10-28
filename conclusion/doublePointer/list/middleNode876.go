package list

// 876. 链表的中间结点
// https://leetcode.cn/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == nil {
			break
		}
	}
	return slow
}
