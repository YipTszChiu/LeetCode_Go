package _00HotTopic

// 21. 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 初始化一个头结点
	head := &ListNode{0, nil}
	// 构建两个指针
	node1, node2, node3 := list1, list2, head
	// 遍历两个升序链表
	for node1 != nil && node2 != nil {
		// 合并的链表指向更小的一个节点，并且合并链表的指针后移一位
		if node1.Val <= node2.Val {
			node3.Next = node1
			node1 = node1.Next
			node3 = node3.Next
		} else {
			node3.Next = node2
			node2 = node2.Next
			node3 = node3.Next
		}
	}
	// 循环结束检查 list1, list2 是否遍历完全
	if node1 != nil {
		for node1 != nil {
			node3.Next = node1
			node1 = node1.Next
			node3 = node3.Next
		}
	}
	if node2 != nil {
		for node2 != nil {
			node3.Next = node2
			node2 = node2.Next
			node3 = node3.Next
		}
	}
	return head.Next
}
