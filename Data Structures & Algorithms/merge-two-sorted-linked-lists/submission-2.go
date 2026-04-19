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
			penis.Next = l
			l = l.Next
		} else if r.Val <= l.Val {
			penis.Next = r
			r = r.Next
		}
		penis = penis.Next
	}

	for l != nil {
		penis.Next = l
		l = l.Next
		penis = penis.Next
	}

	for r != nil {
		penis.Next = r
		r = r.Next
		penis = penis.Next
	}

	return dummy.Next
}
