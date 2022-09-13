package _00HotTopic

// 234.回文链表
// https://leetcode.cn/problems/palindrome-linked-list/?favorite=2cktkvj

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 使用切片保存 val 后使用双指针法
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isPalindrome(head *ListNode) bool {
	// 遍历list保存val
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	// 从中间位置使用双指针遍历数组
	n := len(vals)
	result := true
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			result = false
		}
	}
	return result
}

// 快慢指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 将链表反转
func reverseList234(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 第一个指针返回原链表中间位置
	firstHalfEnd := endOfFirstHalf(head)
	// 将原链表后半部分翻转
	secondHalfStart := reverseList234(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList234(secondHalfStart)
	return result
}
