package list

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 当快慢指针相遇时，假如 slow 走了 k 步，则 fast 走了 2k
	// 设头节点到环入口距离 a，环入口到相遇点距离 b，相遇点继续走 c 回到入口
	// 联立k = a + b; 2k = a + b + n(b+c)可以得到 k = n(b+c)
	// 联立 k = a+b 和 k = n(b+c)可以得到a = n(b+c)-b = (n-1)(b+c) + c
	// 根据这条公式，将slow放回头节点，fast当前在相遇点，两者同速前进，当slow走到入口点，fast从相遇点回到入口
	slow, fast := head, head
	// fast 以 两倍与 slow 的步长前进
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 循环结束判断是否有环
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 根据推导结果，两个指针同速率前进
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
