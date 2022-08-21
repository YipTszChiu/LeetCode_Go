package _00HotTopic

// 141. 环形链表
// https://leetcode.cn/problems/linked-list-cycle/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法一：哈希表
// 时间复杂度：O(n) 取决于链表中
func hasCycle(head *ListNode) bool {
	// 使用一个哈希表记录节点是否已经存在
	nodeSet := make(map[*ListNode]bool)
	// 遍历 List
	for node := head; node != nil; node = node.Next {
		if nodeSet[node] == false {
			nodeSet[node] = true
		} else {
			return true
		}
	}
	return false
}

// 方法二：快慢指针
// 时间复杂度:O(n) 取决于链表
// 空间复杂度:O(1)
func hasCycle2(head *ListNode) bool {
	// 思路：使用两个指针，遍历时一个步长为1，另一个步长为2，如果有环则这两个指针必定相遇
	// 判空
	if head == nil || head.Next == nil {
		return false
	}
	// 初始化两个指针
	slow, fast := head, head.Next
	// 持续遍历链表直至两个指针相等
	for slow != fast {
		// 遇到空指针必然不成环
		if fast == nil || fast.Next == nil {
			return false
		}
		// 快慢指针步长不同
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
