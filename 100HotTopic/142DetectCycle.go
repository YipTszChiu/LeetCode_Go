package _00HotTopic

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 哈希表法
// 时间复杂度：O(n) 取决于链表长度
// 空间复杂度：O(n) 取决于链表长度
func detectCycle(head *ListNode) *ListNode {
	// 使用哈希表记录遍历过的节点
	nodeSet := make(map[*ListNode]bool)
	// 遍历链表
	for node := head; node != nil; node = node.Next {
		if nodeSet[node] == false {
			nodeSet[node] = true
		} else {
			return node
		}
	}
	return nil
}

// 快慢指针
func detectCycle2(head *ListNode) *ListNode {
	// 根据快慢指针法，如果有环，两个指针必定相遇于一点
	// 此时，fast 走过的路程 = a + n(b+c) + b = a + (n+1)b + nc
	// fast 路程永远是 slow 的两倍 a + (n+1)b + nc = 2(a + b) → a =(n-1)b + nc → a = c + (n-1)(b+c)
	// 也就是说，未入环的距离相当于 n-1圈+c
	// 从fast,slow相遇点开始，使用一个指针从head向后移，slow从相遇点向后移，最终会相遇在入环点

	// 初始化快慢指针
	slow, fast := head, head
	// 1. 如果fast为nil，只有一个节点不成环； 2. fast 为 nil 不成环（但这个会在下面跳出）
	for fast != nil {
		// slow 后移一个
		slow = slow.Next
		// fast.Next 为空，不成环，返回 nil
		if fast.Next == nil {
			return nil
		}
		// 此时fast有可能为空，跳出循环时返回nil
		fast = fast.Next.Next
		// 如果fast一直不空，说明右环，slow fast会最终相遇
		if fast == slow {
			// 从相遇点开始，ptr从head开始后移，slow从相遇点后移
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			// ptr 和 slow 的相遇点即为入环点
			return ptr
		}
	}
	return nil
}
