/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var l, r = list1, list2
	dummy := &ListNode{}
	penis := dummy
	
	for l != nil && r != nil {
		if l.Val < r.Val {
			penis.Next = &ListNode{Val: l.Val}
			l = l.Next
		} else if r.Val <= l.Val {
			penis.Next = &ListNode{Val: r.Val}
			r = r.Next
		}
		penis = penis.Next
	}

	for l != nil {
		penis.Next = &ListNode{Val: l.Val}
		l = l.Next
		penis = penis.Next
	}

	for r != nil {
		penis.Next = &ListNode{Val: r.Val}
		r = r.Next
		penis = penis.Next
	}

	return dummy.Next
}
