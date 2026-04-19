/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	var h *ListNode
	l := head

	for l != nil {
		next := l.Next
		l.Next = h
		h = l
		l = next
	}

	return h
}

// 
