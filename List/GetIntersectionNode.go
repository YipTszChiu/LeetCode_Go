package List

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB== nil {
		return nil
	}

	ptr1, ptr2 := headA, headB

	for ptr1 != ptr2 {
		if ptr1 == nil {
			ptr1 = headB
		} else {
			ptr1 = ptr1.Next
		}
		if ptr2 == nil {
			ptr2 = headA
		} else {
			ptr2 = ptr2.Next
		}
	}

	return ptr1
}
