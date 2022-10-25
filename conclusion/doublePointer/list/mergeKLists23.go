package list

// 23. 合并K个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	sortedList := lists[0]
	for i := 1; i < len(lists); i++ {
		sortedList = mergeTwoLists23(sortedList, lists[i])
	}
	return sortedList
}

func mergeTwoLists23(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}

	temp := &ListNode{Val: 0, Next: nil}
	p, p1, p2 := temp, list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return temp.Next
}

// 方法二：类归并排序
func mergeKLists2(lists []*ListNode) *ListNode {
	// 使用 merge 排序 lists 并返回排序后的结果
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	// 左右相加除以二取得中间值
	mid := (left + right) >> 1
	return mergeTwoLists23_2(merge(lists, left, mid), merge(lists, mid+1, right))
}

func mergeTwoLists23_2(list1, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		} else {
			return list1
		}
	}

	head := &ListNode{0, nil}
	tail, p1, p2 := head, list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}

	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}

	return head.Next
}
