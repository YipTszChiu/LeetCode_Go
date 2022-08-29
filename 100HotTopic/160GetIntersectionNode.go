package _00HotTopic

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 思路：A链走完后走B链，B链走完后走A链，如果相交会指向同一个节点，如果不相交会指向 nil
	ptrA, ptrB := headA, headB

	for ptrA != ptrB {
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}

	}

	return ptrA
}
