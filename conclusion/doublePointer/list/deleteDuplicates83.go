package list

// 83. 删除排序链表中的重复元素
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 单指针法
func deleteDuplicates_singlePtr(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	for curr.Next != nil {
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

// 双指针法
func deleteDuplicates(head *ListNode) *ListNode {
	// 链表为空或者链表只有一个元素
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 需要清空后面可能链接的重复元素
	slow.Next = nil

	return head
}
